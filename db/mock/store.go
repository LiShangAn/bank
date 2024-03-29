// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/LiShangAn/bank/db/sqlc (interfaces: IStore)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/LiShangAn/bank/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockIStore is a mock of IStore interface.
type MockIStore struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreMockRecorder
}

// MockIStoreMockRecorder is the mock recorder for MockIStore.
type MockIStoreMockRecorder struct {
	mock *MockIStore
}

// NewMockIStore creates a new mock instance.
func NewMockIStore(ctrl *gomock.Controller) *MockIStore {
	mock := &MockIStore{ctrl: ctrl}
	mock.recorder = &MockIStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStore) EXPECT() *MockIStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockIStore) AddAccountBalance(arg0 context.Context, arg1 db.AddAccountBalanceParams) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockIStoreMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockIStore)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method.
func (m *MockIStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockIStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockIStore)(nil).CreateAccount), arg0, arg1)
}

// CreateEntries mocks base method.
func (m *MockIStore) CreateEntries(arg0 context.Context, arg1 db.CreateEntriesParams) (db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntries", arg0, arg1)
	ret0, _ := ret[0].(db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntries indicates an expected call of CreateEntries.
func (mr *MockIStoreMockRecorder) CreateEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntries", reflect.TypeOf((*MockIStore)(nil).CreateEntries), arg0, arg1)
}

// CreateTransfers mocks base method.
func (m *MockIStore) CreateTransfers(arg0 context.Context, arg1 db.CreateTransfersParams) (db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfers", arg0, arg1)
	ret0, _ := ret[0].(db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfers indicates an expected call of CreateTransfers.
func (mr *MockIStoreMockRecorder) CreateTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfers", reflect.TypeOf((*MockIStore)(nil).CreateTransfers), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockIStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockIStore) DeleteAccount(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockIStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockIStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteEntries mocks base method.
func (m *MockIStore) DeleteEntries(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntries", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntries indicates an expected call of DeleteEntries.
func (mr *MockIStoreMockRecorder) DeleteEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntries", reflect.TypeOf((*MockIStore)(nil).DeleteEntries), arg0, arg1)
}

// DeleteTransfers mocks base method.
func (m *MockIStore) DeleteTransfers(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransfers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransfers indicates an expected call of DeleteTransfers.
func (mr *MockIStoreMockRecorder) DeleteTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransfers", reflect.TypeOf((*MockIStore)(nil).DeleteTransfers), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockIStore) GetAccount(arg0 context.Context, arg1 int32) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockIStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockIStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockIStore) GetAccountForUpdate(arg0 context.Context, arg1 int32) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockIStoreMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockIStore)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetEntries mocks base method.
func (m *MockIStore) GetEntries(arg0 context.Context, arg1 int32) (db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntries", arg0, arg1)
	ret0, _ := ret[0].(db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntries indicates an expected call of GetEntries.
func (mr *MockIStoreMockRecorder) GetEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntries", reflect.TypeOf((*MockIStore)(nil).GetEntries), arg0, arg1)
}

// GetTransfers mocks base method.
func (m *MockIStore) GetTransfers(arg0 context.Context, arg1 int32) (db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfers", arg0, arg1)
	ret0, _ := ret[0].(db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfers indicates an expected call of GetTransfers.
func (mr *MockIStoreMockRecorder) GetTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfers", reflect.TypeOf((*MockIStore)(nil).GetTransfers), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockIStore) GetUser(arg0 context.Context, arg1 string) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIStore)(nil).GetUser), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockIStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockIStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockIStore)(nil).ListAccounts), arg0, arg1)
}

// ListEntries mocks base method.
func (m *MockIStore) ListEntries(arg0 context.Context, arg1 db.ListEntriesParams) ([]db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockIStoreMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockIStore)(nil).ListEntries), arg0, arg1)
}

// ListTransfers mocks base method.
func (m *MockIStore) ListTransfers(arg0 context.Context, arg1 db.ListTransfersParams) ([]db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockIStoreMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockIStore)(nil).ListTransfers), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockIStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockIStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockIStore)(nil).TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockIStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Accounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Accounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockIStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockIStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateEntries mocks base method.
func (m *MockIStore) UpdateEntries(arg0 context.Context, arg1 db.UpdateEntriesParams) (db.Entries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntries", arg0, arg1)
	ret0, _ := ret[0].(db.Entries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntries indicates an expected call of UpdateEntries.
func (mr *MockIStoreMockRecorder) UpdateEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntries", reflect.TypeOf((*MockIStore)(nil).UpdateEntries), arg0, arg1)
}

// UpdateTransfers mocks base method.
func (m *MockIStore) UpdateTransfers(arg0 context.Context, arg1 db.UpdateTransfersParams) (db.Transfers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransfers", arg0, arg1)
	ret0, _ := ret[0].(db.Transfers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransfers indicates an expected call of UpdateTransfers.
func (mr *MockIStoreMockRecorder) UpdateTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransfers", reflect.TypeOf((*MockIStore)(nil).UpdateTransfers), arg0, arg1)
}
