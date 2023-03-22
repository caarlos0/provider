// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	autoscalingv1 "k8s.io/api/autoscaling/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/autoscaling/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// HorizontalPodAutoscalerInterface is an autogenerated mock type for the HorizontalPodAutoscalerInterface type
type HorizontalPodAutoscalerInterface struct {
	mock.Mock
}

type HorizontalPodAutoscalerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *HorizontalPodAutoscalerInterface) EXPECT() *HorizontalPodAutoscalerInterface_Expecter {
	return &HorizontalPodAutoscalerInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, horizontalPodAutoscaler, opts
func (_m *HorizontalPodAutoscalerInterface) Apply(ctx context.Context, horizontalPodAutoscaler *v1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	ret := _m.Called(ctx, horizontalPodAutoscaler, opts)

	var r0 *autoscalingv1.HorizontalPodAutoscaler
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)); ok {
		return rf(ctx, horizontalPodAutoscaler, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) *autoscalingv1.HorizontalPodAutoscaler); ok {
		r0 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscaler)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type HorizontalPodAutoscalerInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - horizontalPodAutoscaler *v1.HorizontalPodAutoscalerApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) Apply(ctx interface{}, horizontalPodAutoscaler interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_Apply_Call {
	return &HorizontalPodAutoscalerInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, horizontalPodAutoscaler, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_Apply_Call) Run(run func(ctx context.Context, horizontalPodAutoscaler *v1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions)) *HorizontalPodAutoscalerInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.HorizontalPodAutoscalerApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Apply_Call) Return(result *autoscalingv1.HorizontalPodAutoscaler, err error) *HorizontalPodAutoscalerInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)) *HorizontalPodAutoscalerInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, horizontalPodAutoscaler, opts
func (_m *HorizontalPodAutoscalerInterface) ApplyStatus(ctx context.Context, horizontalPodAutoscaler *v1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	ret := _m.Called(ctx, horizontalPodAutoscaler, opts)

	var r0 *autoscalingv1.HorizontalPodAutoscaler
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)); ok {
		return rf(ctx, horizontalPodAutoscaler, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) *autoscalingv1.HorizontalPodAutoscaler); ok {
		r0 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscaler)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type HorizontalPodAutoscalerInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - horizontalPodAutoscaler *v1.HorizontalPodAutoscalerApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) ApplyStatus(ctx interface{}, horizontalPodAutoscaler interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_ApplyStatus_Call {
	return &HorizontalPodAutoscalerInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, horizontalPodAutoscaler, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_ApplyStatus_Call) Run(run func(ctx context.Context, horizontalPodAutoscaler *v1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions)) *HorizontalPodAutoscalerInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.HorizontalPodAutoscalerApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_ApplyStatus_Call) Return(result *autoscalingv1.HorizontalPodAutoscaler, err error) *HorizontalPodAutoscalerInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1.HorizontalPodAutoscalerApplyConfiguration, metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)) *HorizontalPodAutoscalerInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, horizontalPodAutoscaler, opts
func (_m *HorizontalPodAutoscalerInterface) Create(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.CreateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	ret := _m.Called(ctx, horizontalPodAutoscaler, opts)

	var r0 *autoscalingv1.HorizontalPodAutoscaler
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.CreateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)); ok {
		return rf(ctx, horizontalPodAutoscaler, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.CreateOptions) *autoscalingv1.HorizontalPodAutoscaler); ok {
		r0 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscaler)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type HorizontalPodAutoscalerInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler
//   - opts metav1.CreateOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) Create(ctx interface{}, horizontalPodAutoscaler interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_Create_Call {
	return &HorizontalPodAutoscalerInterface_Create_Call{Call: _e.mock.On("Create", ctx, horizontalPodAutoscaler, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_Create_Call) Run(run func(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.CreateOptions)) *HorizontalPodAutoscalerInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*autoscalingv1.HorizontalPodAutoscaler), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Create_Call) Return(_a0 *autoscalingv1.HorizontalPodAutoscaler, _a1 error) *HorizontalPodAutoscalerInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Create_Call) RunAndReturn(run func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.CreateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)) *HorizontalPodAutoscalerInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *HorizontalPodAutoscalerInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HorizontalPodAutoscalerInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type HorizontalPodAutoscalerInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_Delete_Call {
	return &HorizontalPodAutoscalerInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *HorizontalPodAutoscalerInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Delete_Call) Return(_a0 error) *HorizontalPodAutoscalerInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *HorizontalPodAutoscalerInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *HorizontalPodAutoscalerInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HorizontalPodAutoscalerInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type HorizontalPodAutoscalerInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *HorizontalPodAutoscalerInterface_DeleteCollection_Call {
	return &HorizontalPodAutoscalerInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *HorizontalPodAutoscalerInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *HorizontalPodAutoscalerInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_DeleteCollection_Call) Return(_a0 error) *HorizontalPodAutoscalerInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *HorizontalPodAutoscalerInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *HorizontalPodAutoscalerInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *autoscalingv1.HorizontalPodAutoscaler
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *autoscalingv1.HorizontalPodAutoscaler); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscaler)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type HorizontalPodAutoscalerInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_Get_Call {
	return &HorizontalPodAutoscalerInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *HorizontalPodAutoscalerInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Get_Call) Return(_a0 *autoscalingv1.HorizontalPodAutoscaler, _a1 error) *HorizontalPodAutoscalerInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)) *HorizontalPodAutoscalerInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *HorizontalPodAutoscalerInterface) List(ctx context.Context, opts metav1.ListOptions) (*autoscalingv1.HorizontalPodAutoscalerList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *autoscalingv1.HorizontalPodAutoscalerList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*autoscalingv1.HorizontalPodAutoscalerList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *autoscalingv1.HorizontalPodAutoscalerList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscalerList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type HorizontalPodAutoscalerInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) List(ctx interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_List_Call {
	return &HorizontalPodAutoscalerInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *HorizontalPodAutoscalerInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_List_Call) Return(_a0 *autoscalingv1.HorizontalPodAutoscalerList, _a1 error) *HorizontalPodAutoscalerInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*autoscalingv1.HorizontalPodAutoscalerList, error)) *HorizontalPodAutoscalerInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *HorizontalPodAutoscalerInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *autoscalingv1.HorizontalPodAutoscaler
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*autoscalingv1.HorizontalPodAutoscaler, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *autoscalingv1.HorizontalPodAutoscaler); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscaler)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type HorizontalPodAutoscalerInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *HorizontalPodAutoscalerInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *HorizontalPodAutoscalerInterface_Patch_Call {
	return &HorizontalPodAutoscalerInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *HorizontalPodAutoscalerInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *HorizontalPodAutoscalerInterface_Patch_Call {
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

func (_c *HorizontalPodAutoscalerInterface_Patch_Call) Return(result *autoscalingv1.HorizontalPodAutoscaler, err error) *HorizontalPodAutoscalerInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*autoscalingv1.HorizontalPodAutoscaler, error)) *HorizontalPodAutoscalerInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, horizontalPodAutoscaler, opts
func (_m *HorizontalPodAutoscalerInterface) Update(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	ret := _m.Called(ctx, horizontalPodAutoscaler, opts)

	var r0 *autoscalingv1.HorizontalPodAutoscaler
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)); ok {
		return rf(ctx, horizontalPodAutoscaler, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) *autoscalingv1.HorizontalPodAutoscaler); ok {
		r0 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscaler)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type HorizontalPodAutoscalerInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler
//   - opts metav1.UpdateOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) Update(ctx interface{}, horizontalPodAutoscaler interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_Update_Call {
	return &HorizontalPodAutoscalerInterface_Update_Call{Call: _e.mock.On("Update", ctx, horizontalPodAutoscaler, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_Update_Call) Run(run func(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.UpdateOptions)) *HorizontalPodAutoscalerInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*autoscalingv1.HorizontalPodAutoscaler), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Update_Call) Return(_a0 *autoscalingv1.HorizontalPodAutoscaler, _a1 error) *HorizontalPodAutoscalerInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Update_Call) RunAndReturn(run func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)) *HorizontalPodAutoscalerInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, horizontalPodAutoscaler, opts
func (_m *HorizontalPodAutoscalerInterface) UpdateStatus(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	ret := _m.Called(ctx, horizontalPodAutoscaler, opts)

	var r0 *autoscalingv1.HorizontalPodAutoscaler
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)); ok {
		return rf(ctx, horizontalPodAutoscaler, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) *autoscalingv1.HorizontalPodAutoscaler); ok {
		r0 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.HorizontalPodAutoscaler)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, horizontalPodAutoscaler, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HorizontalPodAutoscalerInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type HorizontalPodAutoscalerInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler
//   - opts metav1.UpdateOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) UpdateStatus(ctx interface{}, horizontalPodAutoscaler interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_UpdateStatus_Call {
	return &HorizontalPodAutoscalerInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, horizontalPodAutoscaler, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_UpdateStatus_Call) Run(run func(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.UpdateOptions)) *HorizontalPodAutoscalerInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*autoscalingv1.HorizontalPodAutoscaler), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_UpdateStatus_Call) Return(_a0 *autoscalingv1.HorizontalPodAutoscaler, _a1 error) *HorizontalPodAutoscalerInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *autoscalingv1.HorizontalPodAutoscaler, metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error)) *HorizontalPodAutoscalerInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *HorizontalPodAutoscalerInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// HorizontalPodAutoscalerInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type HorizontalPodAutoscalerInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *HorizontalPodAutoscalerInterface_Expecter) Watch(ctx interface{}, opts interface{}) *HorizontalPodAutoscalerInterface_Watch_Call {
	return &HorizontalPodAutoscalerInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *HorizontalPodAutoscalerInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *HorizontalPodAutoscalerInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *HorizontalPodAutoscalerInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HorizontalPodAutoscalerInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *HorizontalPodAutoscalerInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewHorizontalPodAutoscalerInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewHorizontalPodAutoscalerInterface creates a new instance of HorizontalPodAutoscalerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHorizontalPodAutoscalerInterface(t mockConstructorTestingTNewHorizontalPodAutoscalerInterface) *HorizontalPodAutoscalerInterface {
	mock := &HorizontalPodAutoscalerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
