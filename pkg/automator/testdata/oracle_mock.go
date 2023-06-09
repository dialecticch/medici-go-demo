// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/oracles/priceoracle.go

// Package testdata is a generated GoMock package.
package testdata

import (
	big "math/big"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPriceOracle is a mock of PriceOracle interface.
type MockPriceOracle struct {
	ctrl     *gomock.Controller
	recorder *MockPriceOracleMockRecorder
}

// MockPriceOracleMockRecorder is the mock recorder for MockPriceOracle.
type MockPriceOracleMockRecorder struct {
	mock *MockPriceOracle
}

// NewMockPriceOracle creates a new mock instance.
func NewMockPriceOracle(ctrl *gomock.Controller) *MockPriceOracle {
	mock := &MockPriceOracle{ctrl: ctrl}
	mock.recorder = &MockPriceOracleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceOracle) EXPECT() *MockPriceOracleMockRecorder {
	return m.recorder
}

// Decimals mocks base method.
func (m *MockPriceOracle) Decimals() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decimals")
	ret0, _ := ret[0].(int)
	return ret0
}

// Decimals indicates an expected call of Decimals.
func (mr *MockPriceOracleMockRecorder) Decimals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decimals", reflect.TypeOf((*MockPriceOracle)(nil).Decimals))
}

// Price mocks base method.
func (m *MockPriceOracle) Price() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Price")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// Price indicates an expected call of Price.
func (mr *MockPriceOracleMockRecorder) Price() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Price", reflect.TypeOf((*MockPriceOracle)(nil).Price))
}
