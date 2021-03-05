// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/core"
	mock "github.com/stretchr/testify/mock"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

type Plugin_Abort struct {
	*mock.Call
}

func (_m Plugin_Abort) Return(_a0 error) *Plugin_Abort {
	return &Plugin_Abort{Call: _m.Call.Return(_a0)}
}

func (_m *Plugin) OnAbort(ctx context.Context, tCtx core.TaskExecutionContext) *Plugin_Abort {
	c := _m.On("Abort", ctx, tCtx)
	return &Plugin_Abort{Call: c}
}

func (_m *Plugin) OnAbortMatch(matchers ...interface{}) *Plugin_Abort {
	c := _m.On("Abort", matchers...)
	return &Plugin_Abort{Call: c}
}

// Abort provides a mock function with given fields: ctx, tCtx
func (_m *Plugin) Abort(ctx context.Context, tCtx core.TaskExecutionContext) error {
	ret := _m.Called(ctx, tCtx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionContext) error); ok {
		r0 = rf(ctx, tCtx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Plugin_Finalize struct {
	*mock.Call
}

func (_m Plugin_Finalize) Return(_a0 error) *Plugin_Finalize {
	return &Plugin_Finalize{Call: _m.Call.Return(_a0)}
}

func (_m *Plugin) OnFinalize(ctx context.Context, tCtx core.TaskExecutionContext) *Plugin_Finalize {
	c := _m.On("Finalize", ctx, tCtx)
	return &Plugin_Finalize{Call: c}
}

func (_m *Plugin) OnFinalizeMatch(matchers ...interface{}) *Plugin_Finalize {
	c := _m.On("Finalize", matchers...)
	return &Plugin_Finalize{Call: c}
}

// Finalize provides a mock function with given fields: ctx, tCtx
func (_m *Plugin) Finalize(ctx context.Context, tCtx core.TaskExecutionContext) error {
	ret := _m.Called(ctx, tCtx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionContext) error); ok {
		r0 = rf(ctx, tCtx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Plugin_GetID struct {
	*mock.Call
}

func (_m Plugin_GetID) Return(_a0 string) *Plugin_GetID {
	return &Plugin_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *Plugin) OnGetID() *Plugin_GetID {
	c := _m.On("GetID")
	return &Plugin_GetID{Call: c}
}

func (_m *Plugin) OnGetIDMatch(matchers ...interface{}) *Plugin_GetID {
	c := _m.On("GetID", matchers...)
	return &Plugin_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *Plugin) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type Plugin_GetProperties struct {
	*mock.Call
}

func (_m Plugin_GetProperties) Return(_a0 core.PluginProperties) *Plugin_GetProperties {
	return &Plugin_GetProperties{Call: _m.Call.Return(_a0)}
}

func (_m *Plugin) OnGetProperties() *Plugin_GetProperties {
	c := _m.On("GetProperties")
	return &Plugin_GetProperties{Call: c}
}

func (_m *Plugin) OnGetPropertiesMatch(matchers ...interface{}) *Plugin_GetProperties {
	c := _m.On("GetProperties", matchers...)
	return &Plugin_GetProperties{Call: c}
}

// GetProperties provides a mock function with given fields:
func (_m *Plugin) GetProperties() core.PluginProperties {
	ret := _m.Called()

	var r0 core.PluginProperties
	if rf, ok := ret.Get(0).(func() core.PluginProperties); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(core.PluginProperties)
	}

	return r0
}

type Plugin_Handle struct {
	*mock.Call
}

func (_m Plugin_Handle) Return(_a0 core.Transition, _a1 error) *Plugin_Handle {
	return &Plugin_Handle{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Plugin) OnHandle(ctx context.Context, tCtx core.TaskExecutionContext) *Plugin_Handle {
	c := _m.On("Handle", ctx, tCtx)
	return &Plugin_Handle{Call: c}
}

func (_m *Plugin) OnHandleMatch(matchers ...interface{}) *Plugin_Handle {
	c := _m.On("Handle", matchers...)
	return &Plugin_Handle{Call: c}
}

// Handle provides a mock function with given fields: ctx, tCtx
func (_m *Plugin) Handle(ctx context.Context, tCtx core.TaskExecutionContext) (core.Transition, error) {
	ret := _m.Called(ctx, tCtx)

	var r0 core.Transition
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionContext) core.Transition); ok {
		r0 = rf(ctx, tCtx)
	} else {
		r0 = ret.Get(0).(core.Transition)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.TaskExecutionContext) error); ok {
		r1 = rf(ctx, tCtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
