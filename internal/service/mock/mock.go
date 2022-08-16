// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPayments is a mock of Payments interface.
type MockPayments struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentsMockRecorder
}

// MockPaymentsMockRecorder is the mock recorder for MockPayments.
type MockPaymentsMockRecorder struct {
	mock *MockPayments
}

// NewMockPayments creates a new mock instance.
func NewMockPayments(ctrl *gomock.Controller) *MockPayments {
	mock := &MockPayments{ctrl: ctrl}
	mock.recorder = &MockPaymentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPayments) EXPECT() *MockPaymentsMockRecorder {
	return m.recorder
}

// ProductSubscription mocks base method.
func (m *MockPayments) ProductSubscription(productId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProductSubscription", productId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductSubscription indicates an expected call of ProductSubscription.
func (mr *MockPaymentsMockRecorder) ProductSubscription(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductSubscription", reflect.TypeOf((*MockPayments)(nil).ProductSubscription), productId)
}