// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/elliot-zen/microservices/order/internal/application/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// APIPort is an autogenerated mock type for the APIPort type
type APIPort struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *APIPort) Get(ctx context.Context, id int64) (domain.Order, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.Order, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Order); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlaceOrder provides a mock function with given fields: ctx, order
func (_m *APIPort) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for PlaceOrder")
	}

	var r0 domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Order) (domain.Order, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Order) domain.Order); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Order) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAPIPort creates a new instance of APIPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *APIPort {
	mock := &APIPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
