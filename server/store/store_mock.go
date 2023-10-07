// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccountHolder mocks base method.
func (m *MockStore) CreateAccountHolder(ctx context.Context, ah AccountHolder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccountHolder", ctx, ah)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccountHolder indicates an expected call of CreateAccountHolder.
func (mr *MockStoreMockRecorder) CreateAccountHolder(ctx, ah interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountHolder", reflect.TypeOf((*MockStore)(nil).CreateAccountHolder), ctx, ah)
}

// CreateEphemeralAccount mocks base method.
func (m *MockStore) CreateEphemeralAccount(ctx context.Context, ea EphemeralAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEphemeralAccount", ctx, ea)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEphemeralAccount indicates an expected call of CreateEphemeralAccount.
func (mr *MockStoreMockRecorder) CreateEphemeralAccount(ctx, ea interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEphemeralAccount", reflect.TypeOf((*MockStore)(nil).CreateEphemeralAccount), ctx, ea)
}

// FindDormantAccount mocks base method.
func (m *MockStore) FindDormantAccount(ctx context.Context) (*Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDormantAccount", ctx)
	ret0, _ := ret[0].(*Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDormantAccount indicates an expected call of FindDormantAccount.
func (mr *MockStoreMockRecorder) FindDormantAccount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDormantAccount", reflect.TypeOf((*MockStore)(nil).FindDormantAccount), ctx)
}

// FindEphemeralAccountByAccountID mocks base method.
func (m *MockStore) FindEphemeralAccountByAccountID(ctx context.Context, accountID string) (*EphemeralAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEphemeralAccountByAccountID", ctx, accountID)
	ret0, _ := ret[0].(*EphemeralAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEphemeralAccountByAccountID indicates an expected call of FindEphemeralAccountByAccountID.
func (mr *MockStoreMockRecorder) FindEphemeralAccountByAccountID(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEphemeralAccountByAccountID", reflect.TypeOf((*MockStore)(nil).FindEphemeralAccountByAccountID), ctx, accountID)
}

// GetAccountByID mocks base method.
func (m *MockStore) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", ctx, id)
	ret0, _ := ret[0].(*Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockStoreMockRecorder) GetAccountByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockStore)(nil).GetAccountByID), ctx, id)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(ctx context.Context, ac Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, ac)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(ctx, ac interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), ctx, ac)
}

// MockAccountStore is a mock of AccountStore interface.
type MockAccountStore struct {
	ctrl     *gomock.Controller
	recorder *MockAccountStoreMockRecorder
}

// MockAccountStoreMockRecorder is the mock recorder for MockAccountStore.
type MockAccountStoreMockRecorder struct {
	mock *MockAccountStore
}

// NewMockAccountStore creates a new mock instance.
func NewMockAccountStore(ctrl *gomock.Controller) *MockAccountStore {
	mock := &MockAccountStore{ctrl: ctrl}
	mock.recorder = &MockAccountStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountStore) EXPECT() *MockAccountStoreMockRecorder {
	return m.recorder
}

// CreateAccountHolder mocks base method.
func (m *MockAccountStore) CreateAccountHolder(ctx context.Context, ah AccountHolder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccountHolder", ctx, ah)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccountHolder indicates an expected call of CreateAccountHolder.
func (mr *MockAccountStoreMockRecorder) CreateAccountHolder(ctx, ah interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountHolder", reflect.TypeOf((*MockAccountStore)(nil).CreateAccountHolder), ctx, ah)
}

// CreateEphemeralAccount mocks base method.
func (m *MockAccountStore) CreateEphemeralAccount(ctx context.Context, ea EphemeralAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEphemeralAccount", ctx, ea)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEphemeralAccount indicates an expected call of CreateEphemeralAccount.
func (mr *MockAccountStoreMockRecorder) CreateEphemeralAccount(ctx, ea interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEphemeralAccount", reflect.TypeOf((*MockAccountStore)(nil).CreateEphemeralAccount), ctx, ea)
}

// FindDormantAccount mocks base method.
func (m *MockAccountStore) FindDormantAccount(ctx context.Context) (*Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDormantAccount", ctx)
	ret0, _ := ret[0].(*Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDormantAccount indicates an expected call of FindDormantAccount.
func (mr *MockAccountStoreMockRecorder) FindDormantAccount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDormantAccount", reflect.TypeOf((*MockAccountStore)(nil).FindDormantAccount), ctx)
}

// FindEphemeralAccountByAccountID mocks base method.
func (m *MockAccountStore) FindEphemeralAccountByAccountID(ctx context.Context, accountID string) (*EphemeralAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEphemeralAccountByAccountID", ctx, accountID)
	ret0, _ := ret[0].(*EphemeralAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEphemeralAccountByAccountID indicates an expected call of FindEphemeralAccountByAccountID.
func (mr *MockAccountStoreMockRecorder) FindEphemeralAccountByAccountID(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEphemeralAccountByAccountID", reflect.TypeOf((*MockAccountStore)(nil).FindEphemeralAccountByAccountID), ctx, accountID)
}

// GetAccountByID mocks base method.
func (m *MockAccountStore) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", ctx, id)
	ret0, _ := ret[0].(*Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockAccountStoreMockRecorder) GetAccountByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockAccountStore)(nil).GetAccountByID), ctx, id)
}

// UpdateAccount mocks base method.
func (m *MockAccountStore) UpdateAccount(ctx context.Context, ac Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, ac)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountStoreMockRecorder) UpdateAccount(ctx, ac interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountStore)(nil).UpdateAccount), ctx, ac)
}
