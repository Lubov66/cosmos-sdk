package runtime

import (
	rootstore "cosmossdk.io/store/v2/root"
	"fmt"
	"os"
	"slices"

	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"

	runtimev2 "cosmossdk.io/api/cosmos/app/runtime/v2"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"cosmossdk.io/core/registry"
	"cosmossdk.io/core/server"
	"cosmossdk.io/core/store"
	"cosmossdk.io/core/transaction"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	"cosmossdk.io/log"
	"cosmossdk.io/runtime/v2/services"
	"cosmossdk.io/server/v2/stf"
)

var (
	_ appmodulev2.AppModule = appModule[transaction.Tx]{}
	_ hasServicesV1         = appModule[transaction.Tx]{}
)

type appModule[T transaction.Tx] struct {
	app *App[T]
}

func (m appModule[T]) IsOnePerModuleType() {}
func (m appModule[T]) IsAppModule()        {}

func (m appModule[T]) RegisterServices(registrar grpc.ServiceRegistrar) error {
	autoCliQueryService, err := services.NewAutoCLIQueryService(m.app.moduleManager.modules)
	if err != nil {
		return err
	}

	autocliv1.RegisterQueryServer(registrar, autoCliQueryService)

	reflectionSvc, err := services.NewReflectionService()
	if err != nil {
		return err
	}
	reflectionv1.RegisterReflectionServiceServer(registrar, reflectionSvc)

	return nil
}

func (m appModule[T]) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: appv1alpha1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Config",
					Short:     "Query the current app config",
				},
			},
			SubCommands: map[string]*autocliv1.ServiceCommandDescriptor{
				"autocli": {
					Service: autocliv1.Query_ServiceDesc.ServiceName,
					RpcCommandOptions: []*autocliv1.RpcCommandOptions{
						{
							RpcMethod: "AppOptions",
							Short:     "Query the custom autocli options",
						},
					},
				},
				"reflection": {
					Service: reflectionv1.ReflectionService_ServiceDesc.ServiceName,
					RpcCommandOptions: []*autocliv1.RpcCommandOptions{
						{
							RpcMethod: "FileDescriptors",
							Short:     "Query the app's protobuf file descriptors",
						},
					},
				},
			},
		},
	}
}

func init() {
	appconfig.Register(&runtimev2.Module{},
		appconfig.Provide(
			ProvideAppBuilder[transaction.Tx],
			ProvideEnvironment[transaction.Tx],
			ProvideModuleManager[transaction.Tx],
			ProvideStoreBuilder,
		),
		appconfig.Invoke(SetupAppBuilder),
	)
}

func ProvideAppBuilder[T transaction.Tx](
	interfaceRegistrar registry.InterfaceRegistrar,
	amino registry.AminoRegistrar,
) (
	*AppBuilder[T],
	*stf.MsgRouterBuilder,
	appmodulev2.AppModule,
	protodesc.Resolver,
	protoregistry.MessageTypeResolver,
) {
	protoFiles := proto.HybridResolver
	protoTypes := protoregistry.GlobalTypes

	// At startup, check that all proto annotations are correct.
	if err := validateProtoAnnotations(protoFiles); err != nil {
		// Once we switch to using protoreflect-based ante handlers, we might
		// want to panic here instead of logging a warning.
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}

	msgRouterBuilder := stf.NewMsgRouterBuilder()
	app := &App[T]{
		storeKeys:               nil,
		interfaceRegistrar:      interfaceRegistrar,
		amino:                   amino,
		msgRouterBuilder:        msgRouterBuilder,
		queryRouterBuilder:      stf.NewMsgRouterBuilder(), // TODO dedicated query router
		GRPCMethodsToMessageMap: map[string]func() proto.Message{},
	}
	appBuilder := &AppBuilder[T]{app: app}

	return appBuilder, msgRouterBuilder, appModule[T]{app}, protoFiles, protoTypes
}

type AppInputs struct {
	depinject.In

	Config             *runtimev2.Module
	AppBuilder         *AppBuilder[transaction.Tx]
	ModuleManager      *MM[transaction.Tx]
	InterfaceRegistrar registry.InterfaceRegistrar
	LegacyAmino        registry.AminoRegistrar
	Logger             log.Logger
	StoreBuilder       *StoreBuilder
	DynamicConfig      server.DynamicConfig `optional:"true"` // can be nil in client wiring
	StoreOptions       *rootstore.Options   `optional:"true"` // if unset defaults will be used
}

func SetupAppBuilder(inputs AppInputs) {
	app := inputs.AppBuilder.app
	app.config = inputs.Config
	app.logger = inputs.Logger
	app.moduleManager = inputs.ModuleManager
	app.moduleManager.RegisterInterfaces(inputs.InterfaceRegistrar)
	app.moduleManager.RegisterLegacyAminoCodec(inputs.LegacyAmino)

	if inputs.DynamicConfig == nil {
		return
	}
	storeOptions := rootstore.DefaultStoreOptions()
	if inputs.StoreOptions != nil {
		storeOptions = *inputs.StoreOptions
	}
	var err error
	app.db, err = inputs.StoreBuilder.Build(inputs.Logger, app.storeKeys, inputs.DynamicConfig, storeOptions)
	if err != nil {
		panic(err)
	}
}

func ProvideModuleManager[T transaction.Tx](
	logger log.Logger,
	config *runtimev2.Module,
	modules map[string]appmodulev2.AppModule,
) *MM[T] {
	return NewModuleManager[T](logger, config, modules)
}

// ProvideEnvironment provides the environment for keeper modules, while maintaining backward compatibility and provide services directly as well.
func ProvideEnvironment[T transaction.Tx](
	logger log.Logger,
	config *runtimev2.Module,
	key depinject.ModuleKey,
	appBuilder *AppBuilder[T],
) (
	appmodulev2.Environment,
	store.KVStoreService,
	store.MemoryStoreService,
) {
	var (
		kvService    store.KVStoreService     = failingStoreService{}
		memKvService store.MemoryStoreService = failingStoreService{}
	)

	// skips modules that have no store
	if !slices.Contains(config.SkipStoreKeys, key.Name()) {
		var kvStoreKey string
		storeKeyOverride := storeKeyOverride(config, key.Name())
		if storeKeyOverride != nil {
			kvStoreKey = storeKeyOverride.KvStoreKey
		} else {
			kvStoreKey = key.Name()
		}

		registerStoreKey(appBuilder, kvStoreKey)
		kvService = stf.NewKVStoreService([]byte(kvStoreKey))

		memStoreKey := fmt.Sprintf("memory:%s", key.Name())
		registerStoreKey(appBuilder, memStoreKey)
		memKvService = stf.NewMemoryStoreService([]byte(memStoreKey))
	}

	env := appmodulev2.Environment{
		Logger:             logger,
		BranchService:      stf.BranchService{},
		EventService:       stf.NewEventService(),
		GasService:         stf.NewGasMeterService(),
		HeaderService:      stf.HeaderService{},
		QueryRouterService: stf.NewQueryRouterService(),
		MsgRouterService:   stf.NewMsgRouterService([]byte(key.Name())),
		TransactionService: services.NewContextAwareTransactionService(),
		KVStoreService:     kvService,
		MemStoreService:    memKvService,
	}

	return env, kvService, memKvService
}

func registerStoreKey[T transaction.Tx](builder *AppBuilder[T], key string) {
	builder.app.storeKeys = append(builder.app.storeKeys, key)
}

func storeKeyOverride(config *runtimev2.Module, moduleName string) *runtimev2.StoreKeyConfig {
	for _, cfg := range config.OverrideStoreKeys {
		if cfg.ModuleName == moduleName {
			return cfg
		}
	}

	return nil
}
