// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// MutatingWebhookConfigurationInterface is an autogenerated mock type for the MutatingWebhookConfigurationInterface type
type MutatingWebhookConfigurationInterface struct {
	mock.Mock
}

type MutatingWebhookConfigurationInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MutatingWebhookConfigurationInterface) EXPECT() *MutatingWebhookConfigurationInterface_Expecter {
	return &MutatingWebhookConfigurationInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, mutatingWebhookConfiguration, opts
func (_m *MutatingWebhookConfigurationInterface) Apply(ctx context.Context, mutatingWebhookConfiguration *v1beta1.MutatingWebhookConfigurationApplyConfiguration, opts v1.ApplyOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, mutatingWebhookConfiguration, opts)

	var r0 *admissionregistrationv1beta1.MutatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.MutatingWebhookConfigurationApplyConfiguration, v1.ApplyOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)); ok {
		return rf(ctx, mutatingWebhookConfiguration, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.MutatingWebhookConfigurationApplyConfiguration, v1.ApplyOptions) *admissionregistrationv1beta1.MutatingWebhookConfiguration); ok {
		r0 = rf(ctx, mutatingWebhookConfiguration, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1beta1.MutatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.MutatingWebhookConfigurationApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, mutatingWebhookConfiguration, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MutatingWebhookConfigurationInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type MutatingWebhookConfigurationInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - mutatingWebhookConfiguration *v1beta1.MutatingWebhookConfigurationApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) Apply(ctx interface{}, mutatingWebhookConfiguration interface{}, opts interface{}) *MutatingWebhookConfigurationInterface_Apply_Call {
	return &MutatingWebhookConfigurationInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, mutatingWebhookConfiguration, opts)}
}

func (_c *MutatingWebhookConfigurationInterface_Apply_Call) Run(run func(ctx context.Context, mutatingWebhookConfiguration *v1beta1.MutatingWebhookConfigurationApplyConfiguration, opts v1.ApplyOptions)) *MutatingWebhookConfigurationInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.MutatingWebhookConfigurationApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Apply_Call) Return(result *admissionregistrationv1beta1.MutatingWebhookConfiguration, err error) *MutatingWebhookConfigurationInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1beta1.MutatingWebhookConfigurationApplyConfiguration, v1.ApplyOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)) *MutatingWebhookConfigurationInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, mutatingWebhookConfiguration, opts
func (_m *MutatingWebhookConfigurationInterface) Create(ctx context.Context, mutatingWebhookConfiguration *admissionregistrationv1beta1.MutatingWebhookConfiguration, opts v1.CreateOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, mutatingWebhookConfiguration, opts)

	var r0 *admissionregistrationv1beta1.MutatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.CreateOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)); ok {
		return rf(ctx, mutatingWebhookConfiguration, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.CreateOptions) *admissionregistrationv1beta1.MutatingWebhookConfiguration); ok {
		r0 = rf(ctx, mutatingWebhookConfiguration, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1beta1.MutatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.CreateOptions) error); ok {
		r1 = rf(ctx, mutatingWebhookConfiguration, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MutatingWebhookConfigurationInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MutatingWebhookConfigurationInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - mutatingWebhookConfiguration *admissionregistrationv1beta1.MutatingWebhookConfiguration
//   - opts v1.CreateOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) Create(ctx interface{}, mutatingWebhookConfiguration interface{}, opts interface{}) *MutatingWebhookConfigurationInterface_Create_Call {
	return &MutatingWebhookConfigurationInterface_Create_Call{Call: _e.mock.On("Create", ctx, mutatingWebhookConfiguration, opts)}
}

func (_c *MutatingWebhookConfigurationInterface_Create_Call) Run(run func(ctx context.Context, mutatingWebhookConfiguration *admissionregistrationv1beta1.MutatingWebhookConfiguration, opts v1.CreateOptions)) *MutatingWebhookConfigurationInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admissionregistrationv1beta1.MutatingWebhookConfiguration), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Create_Call) Return(_a0 *admissionregistrationv1beta1.MutatingWebhookConfiguration, _a1 error) *MutatingWebhookConfigurationInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Create_Call) RunAndReturn(run func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.CreateOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)) *MutatingWebhookConfigurationInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *MutatingWebhookConfigurationInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MutatingWebhookConfigurationInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MutatingWebhookConfigurationInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *MutatingWebhookConfigurationInterface_Delete_Call {
	return &MutatingWebhookConfigurationInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *MutatingWebhookConfigurationInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *MutatingWebhookConfigurationInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Delete_Call) Return(_a0 error) *MutatingWebhookConfigurationInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *MutatingWebhookConfigurationInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *MutatingWebhookConfigurationInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MutatingWebhookConfigurationInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type MutatingWebhookConfigurationInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *MutatingWebhookConfigurationInterface_DeleteCollection_Call {
	return &MutatingWebhookConfigurationInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *MutatingWebhookConfigurationInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *MutatingWebhookConfigurationInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_DeleteCollection_Call) Return(_a0 error) *MutatingWebhookConfigurationInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *MutatingWebhookConfigurationInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *MutatingWebhookConfigurationInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *admissionregistrationv1beta1.MutatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *admissionregistrationv1beta1.MutatingWebhookConfiguration); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1beta1.MutatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MutatingWebhookConfigurationInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MutatingWebhookConfigurationInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *MutatingWebhookConfigurationInterface_Get_Call {
	return &MutatingWebhookConfigurationInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *MutatingWebhookConfigurationInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *MutatingWebhookConfigurationInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Get_Call) Return(_a0 *admissionregistrationv1beta1.MutatingWebhookConfiguration, _a1 error) *MutatingWebhookConfigurationInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)) *MutatingWebhookConfigurationInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *MutatingWebhookConfigurationInterface) List(ctx context.Context, opts v1.ListOptions) (*admissionregistrationv1beta1.MutatingWebhookConfigurationList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *admissionregistrationv1beta1.MutatingWebhookConfigurationList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*admissionregistrationv1beta1.MutatingWebhookConfigurationList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *admissionregistrationv1beta1.MutatingWebhookConfigurationList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1beta1.MutatingWebhookConfigurationList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MutatingWebhookConfigurationInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MutatingWebhookConfigurationInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) List(ctx interface{}, opts interface{}) *MutatingWebhookConfigurationInterface_List_Call {
	return &MutatingWebhookConfigurationInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *MutatingWebhookConfigurationInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *MutatingWebhookConfigurationInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_List_Call) Return(_a0 *admissionregistrationv1beta1.MutatingWebhookConfigurationList, _a1 error) *MutatingWebhookConfigurationInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*admissionregistrationv1beta1.MutatingWebhookConfigurationList, error)) *MutatingWebhookConfigurationInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *MutatingWebhookConfigurationInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admissionregistrationv1beta1.MutatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *admissionregistrationv1beta1.MutatingWebhookConfiguration); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1beta1.MutatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MutatingWebhookConfigurationInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type MutatingWebhookConfigurationInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *MutatingWebhookConfigurationInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *MutatingWebhookConfigurationInterface_Patch_Call {
	return &MutatingWebhookConfigurationInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *MutatingWebhookConfigurationInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *MutatingWebhookConfigurationInterface_Patch_Call {
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

func (_c *MutatingWebhookConfigurationInterface_Patch_Call) Return(result *admissionregistrationv1beta1.MutatingWebhookConfiguration, err error) *MutatingWebhookConfigurationInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)) *MutatingWebhookConfigurationInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, mutatingWebhookConfiguration, opts
func (_m *MutatingWebhookConfigurationInterface) Update(ctx context.Context, mutatingWebhookConfiguration *admissionregistrationv1beta1.MutatingWebhookConfiguration, opts v1.UpdateOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, mutatingWebhookConfiguration, opts)

	var r0 *admissionregistrationv1beta1.MutatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.UpdateOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)); ok {
		return rf(ctx, mutatingWebhookConfiguration, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.UpdateOptions) *admissionregistrationv1beta1.MutatingWebhookConfiguration); ok {
		r0 = rf(ctx, mutatingWebhookConfiguration, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1beta1.MutatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, mutatingWebhookConfiguration, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MutatingWebhookConfigurationInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MutatingWebhookConfigurationInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - mutatingWebhookConfiguration *admissionregistrationv1beta1.MutatingWebhookConfiguration
//   - opts v1.UpdateOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) Update(ctx interface{}, mutatingWebhookConfiguration interface{}, opts interface{}) *MutatingWebhookConfigurationInterface_Update_Call {
	return &MutatingWebhookConfigurationInterface_Update_Call{Call: _e.mock.On("Update", ctx, mutatingWebhookConfiguration, opts)}
}

func (_c *MutatingWebhookConfigurationInterface_Update_Call) Run(run func(ctx context.Context, mutatingWebhookConfiguration *admissionregistrationv1beta1.MutatingWebhookConfiguration, opts v1.UpdateOptions)) *MutatingWebhookConfigurationInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admissionregistrationv1beta1.MutatingWebhookConfiguration), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Update_Call) Return(_a0 *admissionregistrationv1beta1.MutatingWebhookConfiguration, _a1 error) *MutatingWebhookConfigurationInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Update_Call) RunAndReturn(run func(context.Context, *admissionregistrationv1beta1.MutatingWebhookConfiguration, v1.UpdateOptions) (*admissionregistrationv1beta1.MutatingWebhookConfiguration, error)) *MutatingWebhookConfigurationInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *MutatingWebhookConfigurationInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
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

// MutatingWebhookConfigurationInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MutatingWebhookConfigurationInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *MutatingWebhookConfigurationInterface_Expecter) Watch(ctx interface{}, opts interface{}) *MutatingWebhookConfigurationInterface_Watch_Call {
	return &MutatingWebhookConfigurationInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *MutatingWebhookConfigurationInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *MutatingWebhookConfigurationInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *MutatingWebhookConfigurationInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MutatingWebhookConfigurationInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *MutatingWebhookConfigurationInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMutatingWebhookConfigurationInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMutatingWebhookConfigurationInterface creates a new instance of MutatingWebhookConfigurationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMutatingWebhookConfigurationInterface(t mockConstructorTestingTNewMutatingWebhookConfigurationInterface) *MutatingWebhookConfigurationInterface {
	mock := &MutatingWebhookConfigurationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
