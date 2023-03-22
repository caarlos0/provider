// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
)

// NetworkingV1beta1Interface is an autogenerated mock type for the NetworkingV1beta1Interface type
type NetworkingV1beta1Interface struct {
	mock.Mock
}

type NetworkingV1beta1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *NetworkingV1beta1Interface) EXPECT() *NetworkingV1beta1Interface_Expecter {
	return &NetworkingV1beta1Interface_Expecter{mock: &_m.Mock}
}

// IngressClasses provides a mock function with given fields:
func (_m *NetworkingV1beta1Interface) IngressClasses() v1beta1.IngressClassInterface {
	ret := _m.Called()

	var r0 v1beta1.IngressClassInterface
	if rf, ok := ret.Get(0).(func() v1beta1.IngressClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.IngressClassInterface)
		}
	}

	return r0
}

// NetworkingV1beta1Interface_IngressClasses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IngressClasses'
type NetworkingV1beta1Interface_IngressClasses_Call struct {
	*mock.Call
}

// IngressClasses is a helper method to define mock.On call
func (_e *NetworkingV1beta1Interface_Expecter) IngressClasses() *NetworkingV1beta1Interface_IngressClasses_Call {
	return &NetworkingV1beta1Interface_IngressClasses_Call{Call: _e.mock.On("IngressClasses")}
}

func (_c *NetworkingV1beta1Interface_IngressClasses_Call) Run(run func()) *NetworkingV1beta1Interface_IngressClasses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NetworkingV1beta1Interface_IngressClasses_Call) Return(_a0 v1beta1.IngressClassInterface) *NetworkingV1beta1Interface_IngressClasses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NetworkingV1beta1Interface_IngressClasses_Call) RunAndReturn(run func() v1beta1.IngressClassInterface) *NetworkingV1beta1Interface_IngressClasses_Call {
	_c.Call.Return(run)
	return _c
}

// Ingresses provides a mock function with given fields: namespace
func (_m *NetworkingV1beta1Interface) Ingresses(namespace string) v1beta1.IngressInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.IngressInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.IngressInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.IngressInterface)
		}
	}

	return r0
}

// NetworkingV1beta1Interface_Ingresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ingresses'
type NetworkingV1beta1Interface_Ingresses_Call struct {
	*mock.Call
}

// Ingresses is a helper method to define mock.On call
//   - namespace string
func (_e *NetworkingV1beta1Interface_Expecter) Ingresses(namespace interface{}) *NetworkingV1beta1Interface_Ingresses_Call {
	return &NetworkingV1beta1Interface_Ingresses_Call{Call: _e.mock.On("Ingresses", namespace)}
}

func (_c *NetworkingV1beta1Interface_Ingresses_Call) Run(run func(namespace string)) *NetworkingV1beta1Interface_Ingresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *NetworkingV1beta1Interface_Ingresses_Call) Return(_a0 v1beta1.IngressInterface) *NetworkingV1beta1Interface_Ingresses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NetworkingV1beta1Interface_Ingresses_Call) RunAndReturn(run func(string) v1beta1.IngressInterface) *NetworkingV1beta1Interface_Ingresses_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with given fields:
func (_m *NetworkingV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// NetworkingV1beta1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type NetworkingV1beta1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *NetworkingV1beta1Interface_Expecter) RESTClient() *NetworkingV1beta1Interface_RESTClient_Call {
	return &NetworkingV1beta1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *NetworkingV1beta1Interface_RESTClient_Call) Run(run func()) *NetworkingV1beta1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NetworkingV1beta1Interface_RESTClient_Call) Return(_a0 rest.Interface) *NetworkingV1beta1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *NetworkingV1beta1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *NetworkingV1beta1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewNetworkingV1beta1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewNetworkingV1beta1Interface creates a new instance of NetworkingV1beta1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNetworkingV1beta1Interface(t mockConstructorTestingTNewNetworkingV1beta1Interface) *NetworkingV1beta1Interface {
	mock := &NetworkingV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
