// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
)

// RbacV1alpha1Interface is an autogenerated mock type for the RbacV1alpha1Interface type
type RbacV1alpha1Interface struct {
	mock.Mock
}

type RbacV1alpha1Interface_Expecter struct {
	mock *mock.Mock
}

func (_m *RbacV1alpha1Interface) EXPECT() *RbacV1alpha1Interface_Expecter {
	return &RbacV1alpha1Interface_Expecter{mock: &_m.Mock}
}

// ClusterRoleBindings provides a mock function with given fields:
func (_m *RbacV1alpha1Interface) ClusterRoleBindings() v1alpha1.ClusterRoleBindingInterface {
	ret := _m.Called()

	var r0 v1alpha1.ClusterRoleBindingInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.ClusterRoleBindingInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ClusterRoleBindingInterface)
		}
	}

	return r0
}

// RbacV1alpha1Interface_ClusterRoleBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterRoleBindings'
type RbacV1alpha1Interface_ClusterRoleBindings_Call struct {
	*mock.Call
}

// ClusterRoleBindings is a helper method to define mock.On call
func (_e *RbacV1alpha1Interface_Expecter) ClusterRoleBindings() *RbacV1alpha1Interface_ClusterRoleBindings_Call {
	return &RbacV1alpha1Interface_ClusterRoleBindings_Call{Call: _e.mock.On("ClusterRoleBindings")}
}

func (_c *RbacV1alpha1Interface_ClusterRoleBindings_Call) Run(run func()) *RbacV1alpha1Interface_ClusterRoleBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RbacV1alpha1Interface_ClusterRoleBindings_Call) Return(_a0 v1alpha1.ClusterRoleBindingInterface) *RbacV1alpha1Interface_ClusterRoleBindings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RbacV1alpha1Interface_ClusterRoleBindings_Call) RunAndReturn(run func() v1alpha1.ClusterRoleBindingInterface) *RbacV1alpha1Interface_ClusterRoleBindings_Call {
	_c.Call.Return(run)
	return _c
}

// ClusterRoles provides a mock function with given fields:
func (_m *RbacV1alpha1Interface) ClusterRoles() v1alpha1.ClusterRoleInterface {
	ret := _m.Called()

	var r0 v1alpha1.ClusterRoleInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.ClusterRoleInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ClusterRoleInterface)
		}
	}

	return r0
}

// RbacV1alpha1Interface_ClusterRoles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterRoles'
type RbacV1alpha1Interface_ClusterRoles_Call struct {
	*mock.Call
}

// ClusterRoles is a helper method to define mock.On call
func (_e *RbacV1alpha1Interface_Expecter) ClusterRoles() *RbacV1alpha1Interface_ClusterRoles_Call {
	return &RbacV1alpha1Interface_ClusterRoles_Call{Call: _e.mock.On("ClusterRoles")}
}

func (_c *RbacV1alpha1Interface_ClusterRoles_Call) Run(run func()) *RbacV1alpha1Interface_ClusterRoles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RbacV1alpha1Interface_ClusterRoles_Call) Return(_a0 v1alpha1.ClusterRoleInterface) *RbacV1alpha1Interface_ClusterRoles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RbacV1alpha1Interface_ClusterRoles_Call) RunAndReturn(run func() v1alpha1.ClusterRoleInterface) *RbacV1alpha1Interface_ClusterRoles_Call {
	_c.Call.Return(run)
	return _c
}

// RESTClient provides a mock function with given fields:
func (_m *RbacV1alpha1Interface) RESTClient() rest.Interface {
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

// RbacV1alpha1Interface_RESTClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RESTClient'
type RbacV1alpha1Interface_RESTClient_Call struct {
	*mock.Call
}

// RESTClient is a helper method to define mock.On call
func (_e *RbacV1alpha1Interface_Expecter) RESTClient() *RbacV1alpha1Interface_RESTClient_Call {
	return &RbacV1alpha1Interface_RESTClient_Call{Call: _e.mock.On("RESTClient")}
}

func (_c *RbacV1alpha1Interface_RESTClient_Call) Run(run func()) *RbacV1alpha1Interface_RESTClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RbacV1alpha1Interface_RESTClient_Call) Return(_a0 rest.Interface) *RbacV1alpha1Interface_RESTClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RbacV1alpha1Interface_RESTClient_Call) RunAndReturn(run func() rest.Interface) *RbacV1alpha1Interface_RESTClient_Call {
	_c.Call.Return(run)
	return _c
}

// RoleBindings provides a mock function with given fields: namespace
func (_m *RbacV1alpha1Interface) RoleBindings(namespace string) v1alpha1.RoleBindingInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.RoleBindingInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.RoleBindingInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.RoleBindingInterface)
		}
	}

	return r0
}

// RbacV1alpha1Interface_RoleBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RoleBindings'
type RbacV1alpha1Interface_RoleBindings_Call struct {
	*mock.Call
}

// RoleBindings is a helper method to define mock.On call
//   - namespace string
func (_e *RbacV1alpha1Interface_Expecter) RoleBindings(namespace interface{}) *RbacV1alpha1Interface_RoleBindings_Call {
	return &RbacV1alpha1Interface_RoleBindings_Call{Call: _e.mock.On("RoleBindings", namespace)}
}

func (_c *RbacV1alpha1Interface_RoleBindings_Call) Run(run func(namespace string)) *RbacV1alpha1Interface_RoleBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RbacV1alpha1Interface_RoleBindings_Call) Return(_a0 v1alpha1.RoleBindingInterface) *RbacV1alpha1Interface_RoleBindings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RbacV1alpha1Interface_RoleBindings_Call) RunAndReturn(run func(string) v1alpha1.RoleBindingInterface) *RbacV1alpha1Interface_RoleBindings_Call {
	_c.Call.Return(run)
	return _c
}

// Roles provides a mock function with given fields: namespace
func (_m *RbacV1alpha1Interface) Roles(namespace string) v1alpha1.RoleInterface {
	ret := _m.Called(namespace)

	var r0 v1alpha1.RoleInterface
	if rf, ok := ret.Get(0).(func(string) v1alpha1.RoleInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.RoleInterface)
		}
	}

	return r0
}

// RbacV1alpha1Interface_Roles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Roles'
type RbacV1alpha1Interface_Roles_Call struct {
	*mock.Call
}

// Roles is a helper method to define mock.On call
//   - namespace string
func (_e *RbacV1alpha1Interface_Expecter) Roles(namespace interface{}) *RbacV1alpha1Interface_Roles_Call {
	return &RbacV1alpha1Interface_Roles_Call{Call: _e.mock.On("Roles", namespace)}
}

func (_c *RbacV1alpha1Interface_Roles_Call) Run(run func(namespace string)) *RbacV1alpha1Interface_Roles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RbacV1alpha1Interface_Roles_Call) Return(_a0 v1alpha1.RoleInterface) *RbacV1alpha1Interface_Roles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RbacV1alpha1Interface_Roles_Call) RunAndReturn(run func(string) v1alpha1.RoleInterface) *RbacV1alpha1Interface_Roles_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRbacV1alpha1Interface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRbacV1alpha1Interface creates a new instance of RbacV1alpha1Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRbacV1alpha1Interface(t mockConstructorTestingTNewRbacV1alpha1Interface) *RbacV1alpha1Interface {
	mock := &RbacV1alpha1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
