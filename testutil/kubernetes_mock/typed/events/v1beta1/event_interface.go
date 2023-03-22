// Code generated by mockery v2.23.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	eventsv1beta1 "k8s.io/api/events/v1beta1"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "k8s.io/client-go/applyconfigurations/events/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// EventInterface is an autogenerated mock type for the EventInterface type
type EventInterface struct {
	mock.Mock
}

type EventInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EventInterface) EXPECT() *EventInterface_Expecter {
	return &EventInterface_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: ctx, event, opts
func (_m *EventInterface) Apply(ctx context.Context, event *v1beta1.EventApplyConfiguration, opts v1.ApplyOptions) (*eventsv1beta1.Event, error) {
	ret := _m.Called(ctx, event, opts)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.EventApplyConfiguration, v1.ApplyOptions) (*eventsv1beta1.Event, error)); ok {
		return rf(ctx, event, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.EventApplyConfiguration, v1.ApplyOptions) *eventsv1beta1.Event); ok {
		r0 = rf(ctx, event, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.EventApplyConfiguration, v1.ApplyOptions) error); ok {
		r1 = rf(ctx, event, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type EventInterface_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - ctx context.Context
//   - event *v1beta1.EventApplyConfiguration
//   - opts v1.ApplyOptions
func (_e *EventInterface_Expecter) Apply(ctx interface{}, event interface{}, opts interface{}) *EventInterface_Apply_Call {
	return &EventInterface_Apply_Call{Call: _e.mock.On("Apply", ctx, event, opts)}
}

func (_c *EventInterface_Apply_Call) Run(run func(ctx context.Context, event *v1beta1.EventApplyConfiguration, opts v1.ApplyOptions)) *EventInterface_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1beta1.EventApplyConfiguration), args[2].(v1.ApplyOptions))
	})
	return _c
}

func (_c *EventInterface_Apply_Call) Return(result *eventsv1beta1.Event, err error) *EventInterface_Apply_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *EventInterface_Apply_Call) RunAndReturn(run func(context.Context, *v1beta1.EventApplyConfiguration, v1.ApplyOptions) (*eventsv1beta1.Event, error)) *EventInterface_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, event, opts
func (_m *EventInterface) Create(ctx context.Context, event *eventsv1beta1.Event, opts v1.CreateOptions) (*eventsv1beta1.Event, error) {
	ret := _m.Called(ctx, event, opts)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *eventsv1beta1.Event, v1.CreateOptions) (*eventsv1beta1.Event, error)); ok {
		return rf(ctx, event, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *eventsv1beta1.Event, v1.CreateOptions) *eventsv1beta1.Event); ok {
		r0 = rf(ctx, event, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *eventsv1beta1.Event, v1.CreateOptions) error); ok {
		r1 = rf(ctx, event, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type EventInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - event *eventsv1beta1.Event
//   - opts v1.CreateOptions
func (_e *EventInterface_Expecter) Create(ctx interface{}, event interface{}, opts interface{}) *EventInterface_Create_Call {
	return &EventInterface_Create_Call{Call: _e.mock.On("Create", ctx, event, opts)}
}

func (_c *EventInterface_Create_Call) Run(run func(ctx context.Context, event *eventsv1beta1.Event, opts v1.CreateOptions)) *EventInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*eventsv1beta1.Event), args[2].(v1.CreateOptions))
	})
	return _c
}

func (_c *EventInterface_Create_Call) Return(_a0 *eventsv1beta1.Event, _a1 error) *EventInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_Create_Call) RunAndReturn(run func(context.Context, *eventsv1beta1.Event, v1.CreateOptions) (*eventsv1beta1.Event, error)) *EventInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateWithEventNamespace provides a mock function with given fields: event
func (_m *EventInterface) CreateWithEventNamespace(event *eventsv1beta1.Event) (*eventsv1beta1.Event, error) {
	ret := _m.Called(event)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*eventsv1beta1.Event) (*eventsv1beta1.Event, error)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(*eventsv1beta1.Event) *eventsv1beta1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*eventsv1beta1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_CreateWithEventNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWithEventNamespace'
type EventInterface_CreateWithEventNamespace_Call struct {
	*mock.Call
}

// CreateWithEventNamespace is a helper method to define mock.On call
//   - event *eventsv1beta1.Event
func (_e *EventInterface_Expecter) CreateWithEventNamespace(event interface{}) *EventInterface_CreateWithEventNamespace_Call {
	return &EventInterface_CreateWithEventNamespace_Call{Call: _e.mock.On("CreateWithEventNamespace", event)}
}

func (_c *EventInterface_CreateWithEventNamespace_Call) Run(run func(event *eventsv1beta1.Event)) *EventInterface_CreateWithEventNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*eventsv1beta1.Event))
	})
	return _c
}

func (_c *EventInterface_CreateWithEventNamespace_Call) Return(_a0 *eventsv1beta1.Event, _a1 error) *EventInterface_CreateWithEventNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_CreateWithEventNamespace_Call) RunAndReturn(run func(*eventsv1beta1.Event) (*eventsv1beta1.Event, error)) *EventInterface_CreateWithEventNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *EventInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type EventInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.DeleteOptions
func (_e *EventInterface_Expecter) Delete(ctx interface{}, name interface{}, opts interface{}) *EventInterface_Delete_Call {
	return &EventInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, name, opts)}
}

func (_c *EventInterface_Delete_Call) Run(run func(ctx context.Context, name string, opts v1.DeleteOptions)) *EventInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.DeleteOptions))
	})
	return _c
}

func (_c *EventInterface_Delete_Call) Return(_a0 error) *EventInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventInterface_Delete_Call) RunAndReturn(run func(context.Context, string, v1.DeleteOptions) error) *EventInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *EventInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventInterface_DeleteCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCollection'
type EventInterface_DeleteCollection_Call struct {
	*mock.Call
}

// DeleteCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.DeleteOptions
//   - listOpts v1.ListOptions
func (_e *EventInterface_Expecter) DeleteCollection(ctx interface{}, opts interface{}, listOpts interface{}) *EventInterface_DeleteCollection_Call {
	return &EventInterface_DeleteCollection_Call{Call: _e.mock.On("DeleteCollection", ctx, opts, listOpts)}
}

func (_c *EventInterface_DeleteCollection_Call) Run(run func(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions)) *EventInterface_DeleteCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.DeleteOptions), args[2].(v1.ListOptions))
	})
	return _c
}

func (_c *EventInterface_DeleteCollection_Call) Return(_a0 error) *EventInterface_DeleteCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventInterface_DeleteCollection_Call) RunAndReturn(run func(context.Context, v1.DeleteOptions, v1.ListOptions) error) *EventInterface_DeleteCollection_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *EventInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*eventsv1beta1.Event, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) (*eventsv1beta1.Event, error)); ok {
		return rf(ctx, name, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *eventsv1beta1.Event); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type EventInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts v1.GetOptions
func (_e *EventInterface_Expecter) Get(ctx interface{}, name interface{}, opts interface{}) *EventInterface_Get_Call {
	return &EventInterface_Get_Call{Call: _e.mock.On("Get", ctx, name, opts)}
}

func (_c *EventInterface_Get_Call) Run(run func(ctx context.Context, name string, opts v1.GetOptions)) *EventInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(v1.GetOptions))
	})
	return _c
}

func (_c *EventInterface_Get_Call) Return(_a0 *eventsv1beta1.Event, _a1 error) *EventInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_Get_Call) RunAndReturn(run func(context.Context, string, v1.GetOptions) (*eventsv1beta1.Event, error)) *EventInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, opts
func (_m *EventInterface) List(ctx context.Context, opts v1.ListOptions) (*eventsv1beta1.EventList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *eventsv1beta1.EventList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*eventsv1beta1.EventList, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *eventsv1beta1.EventList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.EventList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type EventInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *EventInterface_Expecter) List(ctx interface{}, opts interface{}) *EventInterface_List_Call {
	return &EventInterface_List_Call{Call: _e.mock.On("List", ctx, opts)}
}

func (_c *EventInterface_List_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *EventInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *EventInterface_List_Call) Return(_a0 *eventsv1beta1.EventList, _a1 error) *EventInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_List_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (*eventsv1beta1.EventList, error)) *EventInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *EventInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*eventsv1beta1.Event, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*eventsv1beta1.Event, error)); ok {
		return rf(ctx, name, pt, data, opts, subresources...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *eventsv1beta1.Event); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type EventInterface_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - pt types.PatchType
//   - data []byte
//   - opts v1.PatchOptions
//   - subresources ...string
func (_e *EventInterface_Expecter) Patch(ctx interface{}, name interface{}, pt interface{}, data interface{}, opts interface{}, subresources ...interface{}) *EventInterface_Patch_Call {
	return &EventInterface_Patch_Call{Call: _e.mock.On("Patch",
		append([]interface{}{ctx, name, pt, data, opts}, subresources...)...)}
}

func (_c *EventInterface_Patch_Call) Run(run func(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string)) *EventInterface_Patch_Call {
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

func (_c *EventInterface_Patch_Call) Return(result *eventsv1beta1.Event, err error) *EventInterface_Patch_Call {
	_c.Call.Return(result, err)
	return _c
}

func (_c *EventInterface_Patch_Call) RunAndReturn(run func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) (*eventsv1beta1.Event, error)) *EventInterface_Patch_Call {
	_c.Call.Return(run)
	return _c
}

// PatchWithEventNamespace provides a mock function with given fields: event, data
func (_m *EventInterface) PatchWithEventNamespace(event *eventsv1beta1.Event, data []byte) (*eventsv1beta1.Event, error) {
	ret := _m.Called(event, data)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*eventsv1beta1.Event, []byte) (*eventsv1beta1.Event, error)); ok {
		return rf(event, data)
	}
	if rf, ok := ret.Get(0).(func(*eventsv1beta1.Event, []byte) *eventsv1beta1.Event); ok {
		r0 = rf(event, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*eventsv1beta1.Event, []byte) error); ok {
		r1 = rf(event, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_PatchWithEventNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchWithEventNamespace'
type EventInterface_PatchWithEventNamespace_Call struct {
	*mock.Call
}

// PatchWithEventNamespace is a helper method to define mock.On call
//   - event *eventsv1beta1.Event
//   - data []byte
func (_e *EventInterface_Expecter) PatchWithEventNamespace(event interface{}, data interface{}) *EventInterface_PatchWithEventNamespace_Call {
	return &EventInterface_PatchWithEventNamespace_Call{Call: _e.mock.On("PatchWithEventNamespace", event, data)}
}

func (_c *EventInterface_PatchWithEventNamespace_Call) Run(run func(event *eventsv1beta1.Event, data []byte)) *EventInterface_PatchWithEventNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*eventsv1beta1.Event), args[1].([]byte))
	})
	return _c
}

func (_c *EventInterface_PatchWithEventNamespace_Call) Return(_a0 *eventsv1beta1.Event, _a1 error) *EventInterface_PatchWithEventNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_PatchWithEventNamespace_Call) RunAndReturn(run func(*eventsv1beta1.Event, []byte) (*eventsv1beta1.Event, error)) *EventInterface_PatchWithEventNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, event, opts
func (_m *EventInterface) Update(ctx context.Context, event *eventsv1beta1.Event, opts v1.UpdateOptions) (*eventsv1beta1.Event, error) {
	ret := _m.Called(ctx, event, opts)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *eventsv1beta1.Event, v1.UpdateOptions) (*eventsv1beta1.Event, error)); ok {
		return rf(ctx, event, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *eventsv1beta1.Event, v1.UpdateOptions) *eventsv1beta1.Event); ok {
		r0 = rf(ctx, event, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *eventsv1beta1.Event, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, event, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type EventInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - event *eventsv1beta1.Event
//   - opts v1.UpdateOptions
func (_e *EventInterface_Expecter) Update(ctx interface{}, event interface{}, opts interface{}) *EventInterface_Update_Call {
	return &EventInterface_Update_Call{Call: _e.mock.On("Update", ctx, event, opts)}
}

func (_c *EventInterface_Update_Call) Run(run func(ctx context.Context, event *eventsv1beta1.Event, opts v1.UpdateOptions)) *EventInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*eventsv1beta1.Event), args[2].(v1.UpdateOptions))
	})
	return _c
}

func (_c *EventInterface_Update_Call) Return(_a0 *eventsv1beta1.Event, _a1 error) *EventInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_Update_Call) RunAndReturn(run func(context.Context, *eventsv1beta1.Event, v1.UpdateOptions) (*eventsv1beta1.Event, error)) *EventInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateWithEventNamespace provides a mock function with given fields: event
func (_m *EventInterface) UpdateWithEventNamespace(event *eventsv1beta1.Event) (*eventsv1beta1.Event, error) {
	ret := _m.Called(event)

	var r0 *eventsv1beta1.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*eventsv1beta1.Event) (*eventsv1beta1.Event, error)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(*eventsv1beta1.Event) *eventsv1beta1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventsv1beta1.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*eventsv1beta1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventInterface_UpdateWithEventNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateWithEventNamespace'
type EventInterface_UpdateWithEventNamespace_Call struct {
	*mock.Call
}

// UpdateWithEventNamespace is a helper method to define mock.On call
//   - event *eventsv1beta1.Event
func (_e *EventInterface_Expecter) UpdateWithEventNamespace(event interface{}) *EventInterface_UpdateWithEventNamespace_Call {
	return &EventInterface_UpdateWithEventNamespace_Call{Call: _e.mock.On("UpdateWithEventNamespace", event)}
}

func (_c *EventInterface_UpdateWithEventNamespace_Call) Run(run func(event *eventsv1beta1.Event)) *EventInterface_UpdateWithEventNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*eventsv1beta1.Event))
	})
	return _c
}

func (_c *EventInterface_UpdateWithEventNamespace_Call) Return(_a0 *eventsv1beta1.Event, _a1 error) *EventInterface_UpdateWithEventNamespace_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_UpdateWithEventNamespace_Call) RunAndReturn(run func(*eventsv1beta1.Event) (*eventsv1beta1.Event, error)) *EventInterface_UpdateWithEventNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *EventInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
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

// EventInterface_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type EventInterface_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - opts v1.ListOptions
func (_e *EventInterface_Expecter) Watch(ctx interface{}, opts interface{}) *EventInterface_Watch_Call {
	return &EventInterface_Watch_Call{Call: _e.mock.On("Watch", ctx, opts)}
}

func (_c *EventInterface_Watch_Call) Run(run func(ctx context.Context, opts v1.ListOptions)) *EventInterface_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.ListOptions))
	})
	return _c
}

func (_c *EventInterface_Watch_Call) Return(_a0 watch.Interface, _a1 error) *EventInterface_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EventInterface_Watch_Call) RunAndReturn(run func(context.Context, v1.ListOptions) (watch.Interface, error)) *EventInterface_Watch_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewEventInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventInterface creates a new instance of EventInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventInterface(t mockConstructorTestingTNewEventInterface) *EventInterface {
	mock := &EventInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
