// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ValidatingWebhookConfigurationInterface is an autogenerated mock type for the ValidatingWebhookConfigurationInterface type
type ValidatingWebhookConfigurationInterface struct {
	mock.Mock
}

type ValidatingWebhookConfigurationInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ValidatingWebhookConfigurationInterface) EXPECT() *ValidatingWebhookConfigurationInterface_Expecter {
	return &ValidatingWebhookConfigurationInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, validatingWebhookConfiguration, opts
func (_m *ValidatingWebhookConfigurationInterface) Apply(ctx context.Context, validatingWebhookConfiguration *v1.ValidatingWebhookConfigurationApplyConfiguration, opts metav1.ApplyOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, validatingWebhookConfiguration, opts)

	var r0 *admissionregistrationv1.ValidatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ValidatingWebhookConfigurationApplyConfiguration, metav1.ApplyOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)); ok {
		return rf(ctx, validatingWebhookConfiguration, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ValidatingWebhookConfigurationApplyConfiguration, metav1.ApplyOptions) *admissionregistrationv1.ValidatingWebhookConfiguration); ok {
		r0 = rf(ctx, validatingWebhookConfiguration, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1.ValidatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ValidatingWebhookConfigurationApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, validatingWebhookConfiguration, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingWebhookConfigurationInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type ValidatingWebhookConfigurationInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - validatingWebhookConfiguration *v1.ValidatingWebhookConfigurationApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) Apply(ctx interface{}, validatingWebhookConfiguration interface{}, opts interface{}) *ValidatingWebhookConfigurationInterface_Apply_Call {
	return &ValidatingWebhookConfigurationInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, validatingWebhookConfiguration, opts)}
}

func (_c *ValidatingWebhookConfigurationInterface_Apply_Call) Run(run func(ctx context.Context, validatingWebhookConfiguration *v1.ValidatingWebhookConfigurationApplyConfiguration, opts metav1.ApplyOptions)) *ValidatingWebhookConfigurationInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ValidatingWebhookConfigurationApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Apply_Call) Return(result *admissionregistrationv1.ValidatingWebhookConfiguration, err error) *ValidatingWebhookConfigurationInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.ValidatingWebhookConfigurationApplyConfiguration, metav1.ApplyOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)) *ValidatingWebhookConfigurationInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, validatingWebhookConfiguration, opts
func (_m *ValidatingWebhookConfigurationInterface) Create(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfiguration, opts metav1.CreateOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, validatingWebhookConfiguration, opts)

	var r0 *admissionregistrationv1.ValidatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.CreateOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)); ok {
		return rf(ctx, validatingWebhookConfiguration, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.CreateOptions) *admissionregistrationv1.ValidatingWebhookConfiguration); ok {
		r0 = rf(ctx, validatingWebhookConfiguration, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1.ValidatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, validatingWebhookConfiguration, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingWebhookConfigurationInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ValidatingWebhookConfigurationInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfiguration
//   - opts metav1.CreateOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) Create(ctx interface{}, validatingWebhookConfiguration interface{}, opts interface{}) *ValidatingWebhookConfigurationInterface_Create_Call {
	return &ValidatingWebhookConfigurationInterface_Create_Call{Call: _e.mock.On("Create", ctx, validatingWebhookConfiguration, opts)}
}

func (_c *ValidatingWebhookConfigurationInterface_Create_Call) Run(run func(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfiguration, opts metav1.CreateOptions)) *ValidatingWebhookConfigurationInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admissionregistrationv1.ValidatingWebhookConfiguration), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Create_Call) Return(_a0 *admissionregistrationv1.ValidatingWebhookConfiguration, _a1 error) *ValidatingWebhookConfigurationInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Create_Call) RunAndReturn(run func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.CreateOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)) *ValidatingWebhookConfigurationInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ValidatingWebhookConfigurationInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidatingWebhookConfigurationInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ValidatingWebhookConfigurationInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *ValidatingWebhookConfigurationInterface_Delete_Call {
	return &ValidatingWebhookConfigurationInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *ValidatingWebhookConfigurationInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *ValidatingWebhookConfigurationInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Delete_Call) Return(_a0 error) *ValidatingWebhookConfigurationInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *ValidatingWebhookConfigurationInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *ValidatingWebhookConfigurationInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidatingWebhookConfigurationInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type ValidatingWebhookConfigurationInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *ValidatingWebhookConfigurationInterface_DeleteCollection_Call {
	return &ValidatingWebhookConfigurationInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *ValidatingWebhookConfigurationInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *ValidatingWebhookConfigurationInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_DeleteCollection_Call) Return(_a0 error) *ValidatingWebhookConfigurationInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *ValidatingWebhookConfigurationInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ValidatingWebhookConfigurationInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *admissionregistrationv1.ValidatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *admissionregistrationv1.ValidatingWebhookConfiguration); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1.ValidatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingWebhookConfigurationInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ValidatingWebhookConfigurationInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *ValidatingWebhookConfigurationInterface_Get_Call {
	return &ValidatingWebhookConfigurationInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *ValidatingWebhookConfigurationInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *ValidatingWebhookConfigurationInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Get_Call) Return(_a0 *admissionregistrationv1.ValidatingWebhookConfiguration, _a1 error) *ValidatingWebhookConfigurationInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)) *ValidatingWebhookConfigurationInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *ValidatingWebhookConfigurationInterface) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1.ValidatingWebhookConfigurationList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *admissionregistrationv1.ValidatingWebhookConfigurationList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*admissionregistrationv1.ValidatingWebhookConfigurationList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *admissionregistrationv1.ValidatingWebhookConfigurationList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1.ValidatingWebhookConfigurationList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingWebhookConfigurationInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ValidatingWebhookConfigurationInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) List(ctx interface{}, opts interface{}) *ValidatingWebhookConfigurationInterface_List_Call {
	return &ValidatingWebhookConfigurationInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *ValidatingWebhookConfigurationInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *ValidatingWebhookConfigurationInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_List_Call) Return(_a0 *admissionregistrationv1.ValidatingWebhookConfigurationList, _a1 error) *ValidatingWebhookConfigurationInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*admissionregistrationv1.ValidatingWebhookConfigurationList, error)) *ValidatingWebhookConfigurationInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ValidatingWebhookConfigurationInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*admissionregistrationv1.ValidatingWebhookConfiguration, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admissionregistrationv1.ValidatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *admissionregistrationv1.ValidatingWebhookConfiguration); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1.ValidatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingWebhookConfigurationInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type ValidatingWebhookConfigurationInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *ValidatingWebhookConfigurationInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *ValidatingWebhookConfigurationInterface_Patch_Call {
	return &ValidatingWebhookConfigurationInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *ValidatingWebhookConfigurationInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *ValidatingWebhookConfigurationInterface_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(types.PatchType), args[3].([]byte), args[4].(metav1.PatchOptions), variadicArgs...)
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Patch_Call) Return(result *admissionregistrationv1.ValidatingWebhookConfiguration, err error) *ValidatingWebhookConfigurationInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)) *ValidatingWebhookConfigurationInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, validatingWebhookConfiguration, opts
func (_m *ValidatingWebhookConfigurationInterface) Update(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfiguration, opts metav1.UpdateOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error) {
	ret := _m.Called(ctx, validatingWebhookConfiguration, opts)

	var r0 *admissionregistrationv1.ValidatingWebhookConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.UpdateOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)); ok {
		return rf(ctx, validatingWebhookConfiguration, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.UpdateOptions) *admissionregistrationv1.ValidatingWebhookConfiguration); ok {
		r0 = rf(ctx, validatingWebhookConfiguration, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admissionregistrationv1.ValidatingWebhookConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, validatingWebhookConfiguration, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingWebhookConfigurationInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ValidatingWebhookConfigurationInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfiguration
//   - opts metav1.UpdateOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) Update(ctx interface{}, validatingWebhookConfiguration interface{}, opts interface{}) *ValidatingWebhookConfigurationInterface_Update_Call {
	return &ValidatingWebhookConfigurationInterface_Update_Call{Call: _e.mock.On("Update", ctx, validatingWebhookConfiguration, opts)}
}

func (_c *ValidatingWebhookConfigurationInterface_Update_Call) Run(run func(ctx context.Context, validatingWebhookConfiguration *admissionregistrationv1.ValidatingWebhookConfiguration, opts metav1.UpdateOptions)) *ValidatingWebhookConfigurationInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admissionregistrationv1.ValidatingWebhookConfiguration), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Update_Call) Return(_a0 *admissionregistrationv1.ValidatingWebhookConfiguration, _a1 error) *ValidatingWebhookConfigurationInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Update_Call) RunAndReturn(run func(context.Context, *admissionregistrationv1.ValidatingWebhookConfiguration, metav1.UpdateOptions) (*admissionregistrationv1.ValidatingWebhookConfiguration, error)) *ValidatingWebhookConfigurationInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ValidatingWebhookConfigurationInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (watch.Interface, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatingWebhookConfigurationInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type ValidatingWebhookConfigurationInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *ValidatingWebhookConfigurationInterface_Expecter) Watch(ctx interface{}, opts interface{}) *ValidatingWebhookConfigurationInterface_Watch_Call {
	return &ValidatingWebhookConfigurationInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *ValidatingWebhookConfigurationInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *ValidatingWebhookConfigurationInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *ValidatingWebhookConfigurationInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ValidatingWebhookConfigurationInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *ValidatingWebhookConfigurationInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewValidatingWebhookConfigurationInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidatingWebhookConfigurationInterface creates a new instance of ValidatingWebhookConfigurationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidatingWebhookConfigurationInterface(t mockConstructorTestingTNewValidatingWebhookConfigurationInterface) *ValidatingWebhookConfigurationInterface {
	mock := &ValidatingWebhookConfigurationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
