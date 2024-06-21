package appdatatest

import (
	"context"
	"fmt"

	"pgregory.net/rapid"

	"cosmossdk.io/schema"
	"cosmossdk.io/schema/appdata"
	schematesting "cosmossdk.io/schema/testing"
	"cosmossdk.io/schema/testing/statesim"
)

type SimulatorOptions struct {
	AppSchema          map[string]schema.ModuleSchema
	Listener           appdata.Listener
	EventAlignedWrites bool
	StateSimOptions    statesim.Options
}

type Simulator struct {
	state        *statesim.App
	options      SimulatorOptions
	blockNum     uint64
	blockDataGen *rapid.Generator[BlockData]
}

type BlockData = []appdata.Packet

func NewSimulator(options SimulatorOptions) *Simulator {
	if options.AppSchema == nil {
		options.AppSchema = schematesting.ExampleAppSchema
	}

	sim := &Simulator{
		state:   statesim.NewApp(options.AppSchema, options.StateSimOptions),
		options: options,
	}

	return sim
}

func (a *Simulator) Initialize() error {
	if f := a.options.Listener.Initialize; f != nil {
		_, err := f(context.Background(), appdata.InitializationData{
			HasEventAlignedWrites: a.options.EventAlignedWrites,
		})
		if err != nil {
			return err
		}
	}

	if f := a.options.Listener.InitializeModuleData; f != nil {
		err := a.state.ScanModuleSchemas(func(moduleName string, moduleSchema schema.ModuleSchema) error {
			return f(appdata.ModuleInitializationData{ModuleName: moduleName, Schema: moduleSchema})
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Simulator) BlockDataGen() *rapid.Generator[BlockData] {
	return a.BlockDataGenN(100)
}

func (a *Simulator) BlockDataGenN(maxUpdatesPerBlock int) *rapid.Generator[BlockData] {
	numUpdatesGen := rapid.IntRange(1, maxUpdatesPerBlock)

	return rapid.Custom(func(t *rapid.T) BlockData {
		var packets BlockData

		updateSet := map[string]bool{}
		// filter out any updates to the same key from this block, otherwise we can end up with weird errors
		updateGen := a.state.UpdateGen().Filter(func(update appdata.ObjectUpdateData) bool {
			_, existing := updateSet[fmt.Sprintf("%s:%v", update.ModuleName, update.Update.Key)]
			return !existing
		})
		numUpdates := numUpdatesGen.Draw(t, "numUpdates")
		for i := 0; i < numUpdates; i++ {
			update := updateGen.Draw(t, fmt.Sprintf("update[%d]", i))
			updateSet[fmt.Sprintf("%s:%v", update.ModuleName, update.Update.Key)] = true
			packets = append(packets, update)
		}

		return packets
	})
}

func (a *Simulator) ProcessBlockData(data BlockData) error {
	a.blockNum++

	if f := a.options.Listener.StartBlock; f != nil {
		err := f(appdata.StartBlockData{Height: a.blockNum})
		if err != nil {
			return err
		}
	}

	for _, packet := range data {
		err := a.options.Listener.SendPacket(packet)
		if err != nil {
			return err
		}

		if updateData, ok := packet.(appdata.ObjectUpdateData); ok {
			err = a.state.ApplyUpdate(updateData.ModuleName, updateData.Update)
			if err != nil {
				return err
			}
		}
	}

	if f := a.options.Listener.Commit; f != nil {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}
