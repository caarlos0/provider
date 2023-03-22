// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import mock "github.com/stretchr/testify/mock"

// ComponentStatusExpansion is an autogenerated mock type for the ComponentStatusExpansion type
type ComponentStatusExpansion struct {
	mock.Mock
}

type ComponentStatusExpansion_Expecter struct {
	mock *mock.Mock
}

func (_m *ComponentStatusExpansion) EXPECT() *ComponentStatusExpansion_Expecter {
	return &ComponentStatusExpansion_Expecter{mock: &_m.Mock}
}

type mockConstructorTestingTNewComponentStatusExpansion interface {
	mock.TestingT
	Cleanup(func())
}

// NewComponentStatusExpansion creates a new instance of ComponentStatusExpansion. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewComponentStatusExpansion(t mockConstructorTestingTNewComponentStatusExpansion) *ComponentStatusExpansion {
	mock := &ComponentStatusExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
