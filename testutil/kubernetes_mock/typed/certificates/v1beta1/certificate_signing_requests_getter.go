// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
)

// CertificateSigningRequestsGetter is an autogenerated mock type for the CertificateSigningRequestsGetter type
type CertificateSigningRequestsGetter struct {
	mock.Mock
}

type CertificateSigningRequestsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *CertificateSigningRequestsGetter) EXPECT() *CertificateSigningRequestsGetter_Expecter {
	return &CertificateSigningRequestsGetter_Expecter{mock: &_m.Mock}
}

// CertificateSigningRequests provides a mock function with given fields:
func (_m *CertificateSigningRequestsGetter) CertificateSigningRequests() v1beta1.CertificateSigningRequestInterface {
	ret := _m.Called()

	var r0 v1beta1.CertificateSigningRequestInterface
	if rf, ok := ret.Get(0).(func() v1beta1.CertificateSigningRequestInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.CertificateSigningRequestInterface)
		}
	}

	return r0
}

// CertificateSigningRequestsGetter_CertificateSigningRequests_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CertificateSigningRequests'
type CertificateSigningRequestsGetter_CertificateSigningRequests_Call struct {
	*mock.Call
}

// CertificateSigningRequests is a helper method to define mock.On call
func (_e *CertificateSigningRequestsGetter_Expecter) CertificateSigningRequests() *CertificateSigningRequestsGetter_CertificateSigningRequests_Call {
	return &CertificateSigningRequestsGetter_CertificateSigningRequests_Call{Call: _e.mock.On("CertificateSigningRequests")}
}

func (_c *CertificateSigningRequestsGetter_CertificateSigningRequests_Call) Run(run func()) *CertificateSigningRequestsGetter_CertificateSigningRequests_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CertificateSigningRequestsGetter_CertificateSigningRequests_Call) Return(_a0 v1beta1.CertificateSigningRequestInterface) *CertificateSigningRequestsGetter_CertificateSigningRequests_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CertificateSigningRequestsGetter_CertificateSigningRequests_Call) RunAndReturn(run func() v1beta1.CertificateSigningRequestInterface) *CertificateSigningRequestsGetter_CertificateSigningRequests_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCertificateSigningRequestsGetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCertificateSigningRequestsGetter creates a new instance of CertificateSigningRequestsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCertificateSigningRequestsGetter(t mockConstructorTestingTNewCertificateSigningRequestsGetter) *CertificateSigningRequestsGetter {
	mock := &CertificateSigningRequestsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
