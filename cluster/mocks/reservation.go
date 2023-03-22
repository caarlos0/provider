// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	typesv1beta3 "github.com/akash-network/akash-api/go/node/types/v1beta3"
	mock "github.com/stretchr/testify/mock"

	v1beta3 "github.com/akash-network/akash-api/go/node/market/v1beta3"
)

// Reservation is an autogenerated mock type for the Reservation type
type Reservation struct {
	mock.Mock
}

type Reservation_Expecter struct {
	mock *mock.Mock
}

func (_m *Reservation) EXPECT() *Reservation_Expecter {
	return &Reservation_Expecter{mock: &_m.Mock}
}

// Allocated provides a mock function with given fields:
func (_m *Reservation) Allocated() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Reservation_Allocated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Allocated'
type Reservation_Allocated_Call struct {
	*mock.Call
}

// Allocated is a helper method to define mock.On call
func (_e *Reservation_Expecter) Allocated() *Reservation_Allocated_Call {
	return &Reservation_Allocated_Call{Call: _e.mock.On("Allocated")}
}

func (_c *Reservation_Allocated_Call) Run(run func()) *Reservation_Allocated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Reservation_Allocated_Call) Return(_a0 bool) *Reservation_Allocated_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reservation_Allocated_Call) RunAndReturn(run func() bool) *Reservation_Allocated_Call {
	_c.Call.Return(run)
	return _c
}

// OrderID provides a mock function with given fields:
func (_m *Reservation) OrderID() v1beta3.OrderID {
	ret := _m.Called()

	var r0 v1beta3.OrderID
	if rf, ok := ret.Get(0).(func() v1beta3.OrderID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1beta3.OrderID)
	}

	return r0
}

// Reservation_OrderID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OrderID'
type Reservation_OrderID_Call struct {
	*mock.Call
}

// OrderID is a helper method to define mock.On call
func (_e *Reservation_Expecter) OrderID() *Reservation_OrderID_Call {
	return &Reservation_OrderID_Call{Call: _e.mock.On("OrderID")}
}

func (_c *Reservation_OrderID_Call) Run(run func()) *Reservation_OrderID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Reservation_OrderID_Call) Return(_a0 v1beta3.OrderID) *Reservation_OrderID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reservation_OrderID_Call) RunAndReturn(run func() v1beta3.OrderID) *Reservation_OrderID_Call {
	_c.Call.Return(run)
	return _c
}

// Resources provides a mock function with given fields:
func (_m *Reservation) Resources() typesv1beta3.ResourceGroup {
	ret := _m.Called()

	var r0 typesv1beta3.ResourceGroup
	if rf, ok := ret.Get(0).(func() typesv1beta3.ResourceGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(typesv1beta3.ResourceGroup)
		}
	}

	return r0
}

// Reservation_Resources_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resources'
type Reservation_Resources_Call struct {
	*mock.Call
}

// Resources is a helper method to define mock.On call
func (_e *Reservation_Expecter) Resources() *Reservation_Resources_Call {
	return &Reservation_Resources_Call{Call: _e.mock.On("Resources")}
}

func (_c *Reservation_Resources_Call) Run(run func()) *Reservation_Resources_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Reservation_Resources_Call) Return(_a0 typesv1beta3.ResourceGroup) *Reservation_Resources_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reservation_Resources_Call) RunAndReturn(run func() typesv1beta3.ResourceGroup) *Reservation_Resources_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewReservation interface {
	mock.TestingT
	Cleanup(func())
}

// NewReservation creates a new instance of Reservation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReservation(t mockConstructorTestingTNewReservation) *Reservation {
	mock := &Reservation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
