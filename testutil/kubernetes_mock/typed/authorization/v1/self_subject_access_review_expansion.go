// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import mock "github.com/stretchr/testify/mock"

// SelfSubjectAccessReviewExpansion is an autogenerated mock type for the SelfSubjectAccessReviewExpansion type
type SelfSubjectAccessReviewExpansion struct {
	mock.Mock
}

type SelfSubjectAccessReviewExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *SelfSubjectAccessReviewExpansion) EXPECT() *SelfSubjectAccessReviewExpansion_Expecter {
	return &SelfSubjectAccessReviewExpansion_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewSelfSubjectAccessReviewExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewSelfSubjectAccessReviewExpansion creates a new instance of SelfSubjectAccessReviewExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSelfSubjectAccessReviewExpansion(t mockConstructorTestingTNewSelfSubjectAccessReviewExpansion) *SelfSubjectAccessReviewExpansion {
	mock := &SelfSubjectAccessReviewExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
