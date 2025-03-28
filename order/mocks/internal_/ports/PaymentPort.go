// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/elliot-zen/microservices/order/internal/application/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// PaymentPort is an autogenerated mock type for the PaymentPort type
type PaymentPort struct {
	mock.Mock
}

// Charge provides a mock function with given fields: _a0
func (_m *PaymentPort) Charge(_a0 *domain.Order) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Charge")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Order) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPaymentPort creates a new instance of PaymentPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentPort {
	mock := &PaymentPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
