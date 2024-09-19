// Code generated by MockGen. DO NOT EDIT.
// Source: x/bank/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	address "cosmossdk.io/core/address"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx context.Context, moduleName string) types.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, moduleName)
	ret0, _ := ret[0].(types.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, moduleName)
}

// GetModuleAccountAndPermissions mocks base method.
func (m *MockAccountKeeper) GetModuleAccountAndPermissions(ctx context.Context, moduleName string) (types.ModuleAccountI, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccountAndPermissions", ctx, moduleName)
	ret0, _ := ret[0].(types.ModuleAccountI)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// GetModuleAccountAndPermissions indicates an expected call of GetModuleAccountAndPermissions.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccountAndPermissions(ctx, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccountAndPermissions", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccountAndPermissions), ctx, moduleName)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// GetModuleAddressAndPermissions mocks base method.
func (m *MockAccountKeeper) GetModuleAddressAndPermissions(moduleName string) (types.AccAddress, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddressAndPermissions", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// GetModuleAddressAndPermissions indicates an expected call of GetModuleAddressAndPermissions.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddressAndPermissions(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddressAndPermissions", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddressAndPermissions), moduleName)
}

// GetModulePermissions mocks base method.
func (m *MockAccountKeeper) GetModulePermissions() map[string]types0.PermissionsForAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModulePermissions")
	ret0, _ := ret[0].(map[string]types0.PermissionsForAddress)
	return ret0
}

// GetModulePermissions indicates an expected call of GetModulePermissions.
func (mr *MockAccountKeeperMockRecorder) GetModulePermissions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModulePermissions", reflect.TypeOf((*MockAccountKeeper)(nil).GetModulePermissions))
}

// HasAccount mocks base method.
func (m *MockAccountKeeper) HasAccount(ctx context.Context, addr types.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAccount", ctx, addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasAccount indicates an expected call of HasAccount.
func (mr *MockAccountKeeperMockRecorder) HasAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAccount", reflect.TypeOf((*MockAccountKeeper)(nil).HasAccount), ctx, addr)
}

// NewAccount mocks base method.
func (m *MockAccountKeeper) NewAccount(arg0 context.Context, arg1 types.AccountI) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0, arg1)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountKeeperMockRecorder) NewAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccount), arg0, arg1)
}

// NewAccountWithAddress mocks base method.
func (m *MockAccountKeeper) NewAccountWithAddress(ctx context.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccountWithAddress", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// NewAccountWithAddress indicates an expected call of NewAccountWithAddress.
func (mr *MockAccountKeeperMockRecorder) NewAccountWithAddress(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccountWithAddress", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccountWithAddress), ctx, addr)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx context.Context, acc types.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// SetModuleAccount mocks base method.
func (m *MockAccountKeeper) SetModuleAccount(ctx context.Context, macc types.ModuleAccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetModuleAccount", ctx, macc)
}

// SetModuleAccount indicates an expected call of SetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) SetModuleAccount(ctx, macc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetModuleAccount), ctx, macc)
}

// ValidatePermissions mocks base method.
func (m *MockAccountKeeper) ValidatePermissions(macc types.ModuleAccountI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePermissions", macc)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePermissions indicates an expected call of ValidatePermissions.
func (mr *MockAccountKeeperMockRecorder) ValidatePermissions(macc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePermissions", reflect.TypeOf((*MockAccountKeeper)(nil).ValidatePermissions), macc)
}
