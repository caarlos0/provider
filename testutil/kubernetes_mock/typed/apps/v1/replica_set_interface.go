// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	appsv1 "k8s.io/api/apps/v1"
	apiautoscalingv1 "k8s.io/api/autoscaling/v1"

	autoscalingv1 "k8s.io/client-go/applyconfigurations/autoscaling/v1"

	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/client-go/applyconfigurations/apps/v1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// ReplicaSetInterface is an autogenerated mock type for the ReplicaSetInterface type
type ReplicaSetInterface struct {
	mock.Mock
}

type ReplicaSetInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ReplicaSetInterface) EXPECT() *ReplicaSetInterface_Expecter {
	return &ReplicaSetInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) Apply(ctx context.Context, replicaSet *v1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *appsv1.ReplicaSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) (*appsv1.ReplicaSet, error)); ok {
		return rf(ctx, replicaSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) *appsv1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type ReplicaSetInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSet *v1.ReplicaSetApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *ReplicaSetInterface_Expecter) Apply(ctx interface{}, replicaSet interface{}, opts interface{}) *ReplicaSetInterface_Apply_Call {
	return &ReplicaSetInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, replicaSet, opts)}
}

func (_c *ReplicaSetInterface_Apply_Call) Run(run func(ctx context.Context, replicaSet *v1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions)) *ReplicaSetInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ReplicaSetApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_Apply_Call) Return(result *appsv1.ReplicaSet, err error) *ReplicaSetInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ReplicaSetInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) (*appsv1.ReplicaSet, error)) *ReplicaSetInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyScale provides a mock function with given fields: ctx, replicaSetName, scale, opts
func (_m *ReplicaSetInterface) ApplyScale(ctx context.Context, replicaSetName string, scale *autoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (*apiautoscalingv1.Scale, error) {
	ret := _m.Called(ctx, replicaSetName, scale, opts)

	var r0 *apiautoscalingv1.Scale
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *autoscalingv1.ScaleApplyConfiguration, metav1.ApplyOptions) (*apiautoscalingv1.Scale, error)); ok {
		return rf(ctx, replicaSetName, scale, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *autoscalingv1.ScaleApplyConfiguration, metav1.ApplyOptions) *apiautoscalingv1.Scale); ok {
		r0 = rf(ctx, replicaSetName, scale, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiautoscalingv1.Scale)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *autoscalingv1.ScaleApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, replicaSetName, scale, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_ApplyScale_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyScale'
type ReplicaSetInterface_ApplyScale_Call struct {
	*mock.Call
}

// ApplyScale is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSetName string
//   - scale *autoscalingv1.ScaleApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *ReplicaSetInterface_Expecter) ApplyScale(ctx interface{}, replicaSetName interface{}, scale interface{}, opts interface{}) *ReplicaSetInterface_ApplyScale_Call {
	return &ReplicaSetInterface_ApplyScale_Call{Call: _e.mock.On("ApplyScale", ctx, replicaSetName, scale, opts)}
}

func (_c *ReplicaSetInterface_ApplyScale_Call) Run(run func(ctx context.Context, replicaSetName string, scale *autoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions)) *ReplicaSetInterface_ApplyScale_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*autoscalingv1.ScaleApplyConfiguration), args[3].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_ApplyScale_Call) Return(_a0 *apiautoscalingv1.Scale, _a1 error) *ReplicaSetInterface_ApplyScale_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_ApplyScale_Call) RunAndReturn(run func(context.Context, string, *autoscalingv1.ScaleApplyConfiguration, metav1.ApplyOptions) (*apiautoscalingv1.Scale, error)) *ReplicaSetInterface_ApplyScale_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyStatus provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) ApplyStatus(ctx context.Context, replicaSet *v1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *appsv1.ReplicaSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) (*appsv1.ReplicaSet, error)); ok {
		return rf(ctx, replicaSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) *appsv1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_ApplyStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyStatus'
type ReplicaSetInterface_ApplyStatus_Call struct {
	*mock.Call
}

// ApplyStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSet *v1.ReplicaSetApplyConfiguration
//   - opts metav1.ApplyOptions
func (_e *ReplicaSetInterface_Expecter) ApplyStatus(ctx interface{}, replicaSet interface{}, opts interface{}) *ReplicaSetInterface_ApplyStatus_Call {
	return &ReplicaSetInterface_ApplyStatus_Call{Call: _e.mock.On("ApplyStatus", ctx, replicaSet, opts)}
}

func (_c *ReplicaSetInterface_ApplyStatus_Call) Run(run func(ctx context.Context, replicaSet *v1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions)) *ReplicaSetInterface_ApplyStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ReplicaSetApplyConfiguration), args[2].(metav1.ApplyOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_ApplyStatus_Call) Return(result *appsv1.ReplicaSet, err error) *ReplicaSetInterface_ApplyStatus_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ReplicaSetInterface_ApplyStatus_Call) RunAndReturn(run func(context.Context, *v1.ReplicaSetApplyConfiguration, metav1.ApplyOptions) (*appsv1.ReplicaSet, error)) *ReplicaSetInterface_ApplyStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) Create(ctx context.Context, replicaSet *appsv1.ReplicaSet, opts metav1.CreateOptions) (*appsv1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *appsv1.ReplicaSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.ReplicaSet, metav1.CreateOptions) (*appsv1.ReplicaSet, error)); ok {
		return rf(ctx, replicaSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.ReplicaSet, metav1.CreateOptions) *appsv1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.ReplicaSet, metav1.CreateOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ReplicaSetInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSet *appsv1.ReplicaSet
//   - opts metav1.CreateOptions
func (_e *ReplicaSetInterface_Expecter) Create(ctx interface{}, replicaSet interface{}, opts interface{}) *ReplicaSetInterface_Create_Call {
	return &ReplicaSetInterface_Create_Call{Call: _e.mock.On("Create", ctx, replicaSet, opts)}
}

func (_c *ReplicaSetInterface_Create_Call) Run(run func(ctx context.Context, replicaSet *appsv1.ReplicaSet, opts metav1.CreateOptions)) *ReplicaSetInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1.ReplicaSet), args[2].(metav1.CreateOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_Create_Call) Return(_a0 *appsv1.ReplicaSet, _a1 error) *ReplicaSetInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_Create_Call) RunAndReturn(run func(context.Context, *appsv1.ReplicaSet, metav1.CreateOptions) (*appsv1.ReplicaSet, error)) *ReplicaSetInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *ReplicaSetInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplicaSetInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ReplicaSetInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.DeleteOptions
func (_e *ReplicaSetInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *ReplicaSetInterface_Delete_Call {
	return &ReplicaSetInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *ReplicaSetInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts metav1.DeleteOptions)) *ReplicaSetInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.DeleteOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_Delete_Call) Return(_a0 error) *ReplicaSetInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReplicaSetInterface_Delete_Call) RunAndReturn(run func(context.Context, string, metav1.DeleteOptions) error) *ReplicaSetInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *ReplicaSetInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplicaSetInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type ReplicaSetInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.DeleteOptions
//   - listOpts metav1.ListOptions
func (_e *ReplicaSetInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *ReplicaSetInterface_DeleteCollection_Call {
	return &ReplicaSetInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *ReplicaSetInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions)) *ReplicaSetInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.DeleteOptions), args[2].(metav1.ListOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_DeleteCollection_Call) Return(_a0 error) *ReplicaSetInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReplicaSetInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, metav1.DeleteOptions, metav1.ListOptions) error) *ReplicaSetInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *ReplicaSetInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*appsv1.ReplicaSet, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *appsv1.ReplicaSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*appsv1.ReplicaSet, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *appsv1.ReplicaSet); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ReplicaSetInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts metav1.GetOptions
func (_e *ReplicaSetInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *ReplicaSetInterface_Get_Call {
	return &ReplicaSetInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *ReplicaSetInterface_Get_Call) Run(run func(ctx context.Context, name string, opts metav1.GetOptions)) *ReplicaSetInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_Get_Call) Return(_a0 *appsv1.ReplicaSet, _a1 error) *ReplicaSetInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_Get_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*appsv1.ReplicaSet, error)) *ReplicaSetInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetScale provides a mock function with given fields: ctx, replicaSetName, options
func (_m *ReplicaSetInterface) GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (*apiautoscalingv1.Scale, error) {
	ret := _m.Called(ctx, replicaSetName, options)

	var r0 *apiautoscalingv1.Scale
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) (*apiautoscalingv1.Scale, error)); ok {
		return rf(ctx, replicaSetName, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, metav1.GetOptions) *apiautoscalingv1.Scale); ok {
		r0 = rf(ctx, replicaSetName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiautoscalingv1.Scale)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, metav1.GetOptions) error); ok {
		r1 = rf(ctx, replicaSetName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_GetScale_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetScale'
type ReplicaSetInterface_GetScale_Call struct {
	*mock.Call
}

// GetScale is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSetName string
//   - options metav1.GetOptions
func (_e *ReplicaSetInterface_Expecter) GetScale(ctx interface{}, replicaSetName interface{}, options interface{}) *ReplicaSetInterface_GetScale_Call {
	return &ReplicaSetInterface_GetScale_Call{Call: _e.mock.On("GetScale", ctx, replicaSetName, options)}
}

func (_c *ReplicaSetInterface_GetScale_Call) Run(run func(ctx context.Context, replicaSetName string, options metav1.GetOptions)) *ReplicaSetInterface_GetScale_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(metav1.GetOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_GetScale_Call) Return(_a0 *apiautoscalingv1.Scale, _a1 error) *ReplicaSetInterface_GetScale_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_GetScale_Call) RunAndReturn(run func(context.Context, string, metav1.GetOptions) (*apiautoscalingv1.Scale, error)) *ReplicaSetInterface_GetScale_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *ReplicaSetInterface) List(ctx context.Context, opts metav1.ListOptions) (*appsv1.ReplicaSetList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *appsv1.ReplicaSetList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) (*appsv1.ReplicaSetList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, metav1.ListOptions) *appsv1.ReplicaSetList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSetList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, metav1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ReplicaSetInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *ReplicaSetInterface_Expecter) List(ctx interface{}, opts interface{}) *ReplicaSetInterface_List_Call {
	return &ReplicaSetInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *ReplicaSetInterface_List_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *ReplicaSetInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_List_Call) Return(_a0 *appsv1.ReplicaSetList, _a1 error) *ReplicaSetInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_List_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (*appsv1.ReplicaSetList, error)) *ReplicaSetInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *ReplicaSetInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appsv1.ReplicaSet, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *appsv1.ReplicaSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*appsv1.ReplicaSet, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) *appsv1.ReplicaSet); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type ReplicaSetInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts metav1.PatchOptions
//   - subresources ...string
func (_e *ReplicaSetInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *ReplicaSetInterface_Patch_Call {
	return &ReplicaSetInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *ReplicaSetInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string)) *ReplicaSetInterface_Patch_Call {
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

func (_c *ReplicaSetInterface_Patch_Call) Return(result *appsv1.ReplicaSet, err error) *ReplicaSetInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *ReplicaSetInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, metav1.PatchOptions, ...string) (*appsv1.ReplicaSet, error)) *ReplicaSetInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) Update(ctx context.Context, replicaSet *appsv1.ReplicaSet, opts metav1.UpdateOptions) (*appsv1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *appsv1.ReplicaSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) (*appsv1.ReplicaSet, error)); ok {
		return rf(ctx, replicaSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) *appsv1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ReplicaSetInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSet *appsv1.ReplicaSet
//   - opts metav1.UpdateOptions
func (_e *ReplicaSetInterface_Expecter) Update(ctx interface{}, replicaSet interface{}, opts interface{}) *ReplicaSetInterface_Update_Call {
	return &ReplicaSetInterface_Update_Call{Call: _e.mock.On("Update", ctx, replicaSet, opts)}
}

func (_c *ReplicaSetInterface_Update_Call) Run(run func(ctx context.Context, replicaSet *appsv1.ReplicaSet, opts metav1.UpdateOptions)) *ReplicaSetInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1.ReplicaSet), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_Update_Call) Return(_a0 *appsv1.ReplicaSet, _a1 error) *ReplicaSetInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_Update_Call) RunAndReturn(run func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) (*appsv1.ReplicaSet, error)) *ReplicaSetInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateScale provides a mock function with given fields: ctx, replicaSetName, scale, opts
func (_m *ReplicaSetInterface) UpdateScale(ctx context.Context, replicaSetName string, scale *apiautoscalingv1.Scale, opts metav1.UpdateOptions) (*apiautoscalingv1.Scale, error) {
	ret := _m.Called(ctx, replicaSetName, scale, opts)

	var r0 *apiautoscalingv1.Scale
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *apiautoscalingv1.Scale, metav1.UpdateOptions) (*apiautoscalingv1.Scale, error)); ok {
		return rf(ctx, replicaSetName, scale, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *apiautoscalingv1.Scale, metav1.UpdateOptions) *apiautoscalingv1.Scale); ok {
		r0 = rf(ctx, replicaSetName, scale, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiautoscalingv1.Scale)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *apiautoscalingv1.Scale, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, replicaSetName, scale, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_UpdateScale_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateScale'
type ReplicaSetInterface_UpdateScale_Call struct {
	*mock.Call
}

// UpdateScale is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSetName string
//   - scale *apiautoscalingv1.Scale
//   - opts metav1.UpdateOptions
func (_e *ReplicaSetInterface_Expecter) UpdateScale(ctx interface{}, replicaSetName interface{}, scale interface{}, opts interface{}) *ReplicaSetInterface_UpdateScale_Call {
	return &ReplicaSetInterface_UpdateScale_Call{Call: _e.mock.On("UpdateScale", ctx, replicaSetName, scale, opts)}
}

func (_c *ReplicaSetInterface_UpdateScale_Call) Run(run func(ctx context.Context, replicaSetName string, scale *apiautoscalingv1.Scale, opts metav1.UpdateOptions)) *ReplicaSetInterface_UpdateScale_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*apiautoscalingv1.Scale), args[3].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_UpdateScale_Call) Return(_a0 *apiautoscalingv1.Scale, _a1 error) *ReplicaSetInterface_UpdateScale_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_UpdateScale_Call) RunAndReturn(run func(context.Context, string, *apiautoscalingv1.Scale, metav1.UpdateOptions) (*apiautoscalingv1.Scale, error)) *ReplicaSetInterface_UpdateScale_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: ctx, replicaSet, opts
func (_m *ReplicaSetInterface) UpdateStatus(ctx context.Context, replicaSet *appsv1.ReplicaSet, opts metav1.UpdateOptions) (*appsv1.ReplicaSet, error) {
	ret := _m.Called(ctx, replicaSet, opts)

	var r0 *appsv1.ReplicaSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) (*appsv1.ReplicaSet, error)); ok {
		return rf(ctx, replicaSet, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) *appsv1.ReplicaSet); ok {
		r0 = rf(ctx, replicaSet, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*appsv1.ReplicaSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) error); ok {
		r1 = rf(ctx, replicaSet, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicaSetInterface_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type ReplicaSetInterface_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - replicaSet *appsv1.ReplicaSet
//   - opts metav1.UpdateOptions
func (_e *ReplicaSetInterface_Expecter) UpdateStatus(ctx interface{}, replicaSet interface{}, opts interface{}) *ReplicaSetInterface_UpdateStatus_Call {
	return &ReplicaSetInterface_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", ctx, replicaSet, opts)}
}

func (_c *ReplicaSetInterface_UpdateStatus_Call) Run(run func(ctx context.Context, replicaSet *appsv1.ReplicaSet, opts metav1.UpdateOptions)) *ReplicaSetInterface_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*appsv1.ReplicaSet), args[2].(metav1.UpdateOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_UpdateStatus_Call) Return(_a0 *appsv1.ReplicaSet, _a1 error) *ReplicaSetInterface_UpdateStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_UpdateStatus_Call) RunAndReturn(run func(context.Context, *appsv1.ReplicaSet, metav1.UpdateOptions) (*appsv1.ReplicaSet, error)) *ReplicaSetInterface_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *ReplicaSetInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
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

// ReplicaSetInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type ReplicaSetInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts metav1.ListOptions
func (_e *ReplicaSetInterface_Expecter) Watch(ctx interface{}, opts interface{}) *ReplicaSetInterface_Watch_Call {
	return &ReplicaSetInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *ReplicaSetInterface_Watch_Call) Run(run func(ctx context.Context, opts metav1.ListOptions)) *ReplicaSetInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metav1.ListOptions))
	})
	return _c
}

func (_c *ReplicaSetInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *ReplicaSetInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReplicaSetInterface_Watch_Call) RunAndReturn(run func(context.Context, metav1.ListOptions) (watch.Interface, error)) *ReplicaSetInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewReplicaSetInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewReplicaSetInterface creates a new instance of ReplicaSetInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReplicaSetInterface(t mockConstructorTestingTNewReplicaSetInterface) *ReplicaSetInterface {
	mock := &ReplicaSetInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
