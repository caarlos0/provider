// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
)

// PodSecurityPoliciesGetter is an autogenerated mock type for the PodSecurityPoliciesGetter type
type PodSecurityPoliciesGetter struct {
	mock.Mock
}

type PodSecurityPoliciesGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *PodSecurityPoliciesGetter) EXPECT() *PodSecurityPoliciesGetter_Expecter {
	return &PodSecurityPoliciesGetter_Expecter{mock: &_m.Mock}
}

// PodSecurityPolicies provides a mock function with given fields:
func (_m *PodSecurityPoliciesGetter) PodSecurityPolicies() v1beta1.PodSecurityPolicyInterface {
	ret := _m.Called()

	var r0 v1beta1.PodSecurityPolicyInterface
	if rf, ok := ret.Get(0).(func() v1beta1.PodSecurityPolicyInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.PodSecurityPolicyInterface)
		}
	}

	return r0
}

// PodSecurityPoliciesGetter_PodSecurityPolicies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PodSecurityPolicies'
type PodSecurityPoliciesGetter_PodSecurityPolicies_Call struct {
	*mock.Call
}

// PodSecurityPolicies is a helper method to define mock.On call
func (_e *PodSecurityPoliciesGetter_Expecter) PodSecurityPolicies() *PodSecurityPoliciesGetter_PodSecurityPolicies_Call {
	return &PodSecurityPoliciesGetter_PodSecurityPolicies_Call{Call: _e.mock.On("PodSecurityPolicies")}
}

func (_c *PodSecurityPoliciesGetter_PodSecurityPolicies_Call) Run(run func()) *PodSecurityPoliciesGetter_PodSecurityPolicies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PodSecurityPoliciesGetter_PodSecurityPolicies_Call) Return(_a0 v1beta1.PodSecurityPolicyInterface) *PodSecurityPoliciesGetter_PodSecurityPolicies_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PodSecurityPoliciesGetter_PodSecurityPolicies_Call) RunAndReturn(run func() v1beta1.PodSecurityPolicyInterface) *PodSecurityPoliciesGetter_PodSecurityPolicies_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPodSecurityPoliciesGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewPodSecurityPoliciesGetter creates a new instance of PodSecurityPoliciesGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodSecurityPoliciesGetter(t mockConstructorTestingTNewPodSecurityPoliciesGetter) *PodSecurityPoliciesGetter {
	mock := &PodSecurityPoliciesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
