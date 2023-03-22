// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	appsv1beta2 "k8s.io/api/apps/v1beta2"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta2 "k8s.io/client-go/applyconfigurations/apps/v1beta2"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ControllerRevisionInterface is an autogenerated mock type for the ControllerRevisionInterface type
type ControllerRevisionInterface struct {
	mock.Mock
}

type ControllerRevisionInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ControllerRevisionInterface) EXPECT() *ControllerRevisionInterface_Expecter {
	return &ControllerRevisionInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, controllerRevision, opts
func (_m *ControllerRevisionInterface) Apply(ctx context.Context, controllerRevision *v1beta2.ControllerRevisionApplyConfiguration, opts v1.ApplyOptions) (*appsv1beta2.ControllerRevision, error) {
	ret := _m.Called(ctx, controllerRevision, opts)

	var r0 *appsv1beta2.ControllerRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta2.ControllerRevisionApplyConfiguration, v1.ApplyOptions) (*appsv1beta2.ControllerRevision, error)); ok {
		return rf(ctx, controllerRevision, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta2.ControllerRevisionApplyConfiguration, v1.ApplyOptions) *appsv1beta2.ControllerRevision); ok {
		r0 = rf(ctx, controllerRevision, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta2.ControllerRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta2.ControllerRevisionApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, controllerRevision, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerRevisionInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type ControllerRevisionInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - controllerRevision *v1beta2.ControllerRevisionApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *ControllerRevisionInterface_Expecter) Apply(ctx interface{}, controllerRevision interface{}, opts interface{}) *ControllerRevisionInterface_Apply_Call {
	return &ControllerRevisionInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, controllerRevision, opts)}
}

func (_c *ControllerRevisionInterface_Apply_Call) Run(run func(ctx context.Context, controllerRevision *v1beta2.ControllerRevisionApplyConfiguration, opts v1.ApplyOptions)) *ControllerRevisionInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta2.ControllerRevisionApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_Apply_Call) Return(result *appsv1beta2.ControllerRevision, err error) *ControllerRevisionInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ControllerRevisionInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1beta2.ControllerRevisionApplyConfiguration, v1.ApplyOptions) (*appsv1beta2.ControllerRevision, error)) *ControllerRevisionInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, controllerRevision, opts
func (_m *ControllerRevisionInterface) Create(ctx context.Context, controllerRevision *appsv1beta2.ControllerRevision, opts v1.CreateOptions) (*appsv1beta2.ControllerRevision, error) {
	ret := _m.Called(ctx, controllerRevision, opts)

	var r0 *appsv1beta2.ControllerRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta2.ControllerRevision, v1.CreateOptions) (*appsv1beta2.ControllerRevision, error)); ok {
		return rf(ctx, controllerRevision, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta2.ControllerRevision, v1.CreateOptions) *appsv1beta2.ControllerRevision); ok {
		r0 = rf(ctx, controllerRevision, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta2.ControllerRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1beta2.ControllerRevision, v1.CreateOptions) error); ok {
		r1 = rf(ctx, controllerRevision, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerRevisionInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ControllerRevisionInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - controllerRevision *appsv1beta2.ControllerRevision
//   - opts v1.CreateOptions
func (_e *ControllerRevisionInterface_Expecter) Create(ctx interface{}, controllerRevision interface{}, opts interface{}) *ControllerRevisionInterface_Create_Call {
	return &ControllerRevisionInterface_Create_Call{Call: _e.mock.On("Create", ctx, controllerRevision, opts)}
}

func (_c *ControllerRevisionInterface_Create_Call) Run(run func(ctx context.Context, controllerRevision *appsv1beta2.ControllerRevision, opts v1.CreateOptions)) *ControllerRevisionInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1beta2.ControllerRevision), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_Create_Call) Return(_a0 *appsv1beta2.ControllerRevision, _a1 error) *ControllerRevisionInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerRevisionInterface_Create_Call) RunAndReturn(run func(context.Context, *appsv1beta2.ControllerRevision, v1.CreateOptions) (*appsv1beta2.ControllerRevision, error)) *ControllerRevisionInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ControllerRevisionInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ControllerRevisionInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ControllerRevisionInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *ControllerRevisionInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *ControllerRevisionInterface_Delete_Call {
	return &ControllerRevisionInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *ControllerRevisionInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *ControllerRevisionInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_Delete_Call) Return(_a0 error) *ControllerRevisionInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ControllerRevisionInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *ControllerRevisionInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *ControllerRevisionInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ControllerRevisionInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type ControllerRevisionInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *ControllerRevisionInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *ControllerRevisionInterface_DeleteCollection_Call {
	return &ControllerRevisionInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *ControllerRevisionInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *ControllerRevisionInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_DeleteCollection_Call) Return(_a0 error) *ControllerRevisionInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ControllerRevisionInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *ControllerRevisionInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ControllerRevisionInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*appsv1beta2.ControllerRevision, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *appsv1beta2.ControllerRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*appsv1beta2.ControllerRevision, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *appsv1beta2.ControllerRevision); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta2.ControllerRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerRevisionInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ControllerRevisionInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *ControllerRevisionInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *ControllerRevisionInterface_Get_Call {
	return &ControllerRevisionInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *ControllerRevisionInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *ControllerRevisionInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_Get_Call) Return(_a0 *appsv1beta2.ControllerRevision, _a1 error) *ControllerRevisionInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerRevisionInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*appsv1beta2.ControllerRevision, error)) *ControllerRevisionInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *ControllerRevisionInterface) List(ctx context.Context, opts v1.ListOptions) (*appsv1beta2.ControllerRevisionList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *appsv1beta2.ControllerRevisionList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*appsv1beta2.ControllerRevisionList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *appsv1beta2.ControllerRevisionList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta2.ControllerRevisionList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerRevisionInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ControllerRevisionInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *ControllerRevisionInterface_Expecter) List(ctx interface{}, opts interface{}) *ControllerRevisionInterface_List_Call {
	return &ControllerRevisionInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *ControllerRevisionInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *ControllerRevisionInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_List_Call) Return(_a0 *appsv1beta2.ControllerRevisionList, _a1 error) *ControllerRevisionInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerRevisionInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*appsv1beta2.ControllerRevisionList, error)) *ControllerRevisionInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ControllerRevisionInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*appsv1beta2.ControllerRevision, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *appsv1beta2.ControllerRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*appsv1beta2.ControllerRevision, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *appsv1beta2.ControllerRevision); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta2.ControllerRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerRevisionInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type ControllerRevisionInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *ControllerRevisionInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *ControllerRevisionInterface_Patch_Call {
	return &ControllerRevisionInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *ControllerRevisionInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *ControllerRevisionInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(types.PatchType), args[3].([]byte), args[4].(v1.PatchOptions), variadicArgs...)
	})
	return _c
}

func (_c *ControllerRevisionInterface_Patch_Call) Return(result *appsv1beta2.ControllerRevision, err error) *ControllerRevisionInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ControllerRevisionInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*appsv1beta2.ControllerRevision, error)) *ControllerRevisionInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, controllerRevision, opts
func (_m *ControllerRevisionInterface) Update(ctx context.Context, controllerRevision *appsv1beta2.ControllerRevision, opts v1.UpdateOptions) (*appsv1beta2.ControllerRevision, error) {
	ret := _m.Called(ctx, controllerRevision, opts)

	var r0 *appsv1beta2.ControllerRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta2.ControllerRevision, v1.UpdateOptions) (*appsv1beta2.ControllerRevision, error)); ok {
		return rf(ctx, controllerRevision, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1beta2.ControllerRevision, v1.UpdateOptions) *appsv1beta2.ControllerRevision); ok {
		r0 = rf(ctx, controllerRevision, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1beta2.ControllerRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1beta2.ControllerRevision, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, controllerRevision, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerRevisionInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ControllerRevisionInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - controllerRevision *appsv1beta2.ControllerRevision
//   - opts v1.UpdateOptions
func (_e *ControllerRevisionInterface_Expecter) Update(ctx interface{}, controllerRevision interface{}, opts interface{}) *ControllerRevisionInterface_Update_Call {
	return &ControllerRevisionInterface_Update_Call{Call: _e.mock.On("Update", ctx, controllerRevision, opts)}
}

func (_c *ControllerRevisionInterface_Update_Call) Run(run func(ctx context.Context, controllerRevision *appsv1beta2.ControllerRevision, opts v1.UpdateOptions)) *ControllerRevisionInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1beta2.ControllerRevision), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_Update_Call) Return(_a0 *appsv1beta2.ControllerRevision, _a1 error) *ControllerRevisionInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerRevisionInterface_Update_Call) RunAndReturn(run func(context.Context, *appsv1beta2.ControllerRevision, v1.UpdateOptions) (*appsv1beta2.ControllerRevision, error)) *ControllerRevisionInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ControllerRevisionInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerRevisionInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type ControllerRevisionInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *ControllerRevisionInterface_Expecter) Watch(ctx interface{}, opts interface{}) *ControllerRevisionInterface_Watch_Call {
	return &ControllerRevisionInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *ControllerRevisionInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *ControllerRevisionInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *ControllerRevisionInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *ControllerRevisionInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerRevisionInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *ControllerRevisionInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewControllerRevisionInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewControllerRevisionInterface creates a new instance of ControllerRevisionInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewControllerRevisionInterface(t mockConstructorTestingTNewControllerRevisionInterface) *ControllerRevisionInterface {
	mock := &ControllerRevisionInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
