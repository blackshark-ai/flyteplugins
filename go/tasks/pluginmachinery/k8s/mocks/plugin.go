// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	client "sigs.k8s.io/controller-runtime/pkg/client"

	core "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/core"

	mock "github.com/stretchr/testify/mock"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

type Plugin_BuildIdentityResource struct {
	*mock.Call
}

func (_m Plugin_BuildIdentityResource) Return(_a0 client.Object, _a1 error) *Plugin_BuildIdentityResource {
	return &Plugin_BuildIdentityResource{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Plugin) OnBuildIdentityResource(ctx context.Context, taskCtx core.TaskExecutionMetadata) *Plugin_BuildIdentityResource {
	c := _m.On("BuildIdentityResource", ctx, taskCtx)
	return &Plugin_BuildIdentityResource{Call: c}
}

func (_m *Plugin) OnBuildIdentityResourceMatch(matchers ...interface{}) *Plugin_BuildIdentityResource {
	c := _m.On("BuildIdentityResource", matchers...)
	return &Plugin_BuildIdentityResource{Call: c}
}

// BuildIdentityResource provides a mock function with given fields: ctx, taskCtx
func (_m *Plugin) BuildIdentityResource(ctx context.Context, taskCtx core.TaskExecutionMetadata) (client.Object, error) {
	ret := _m.Called(ctx, taskCtx)

	var r0 client.Object
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionMetadata) client.Object); ok {
		r0 = rf(ctx, taskCtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.TaskExecutionMetadata) error); ok {
		r1 = rf(ctx, taskCtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Plugin_BuildResource struct {
	*mock.Call
}

func (_m Plugin_BuildResource) Return(_a0 client.Object, _a1 error) *Plugin_BuildResource {
	return &Plugin_BuildResource{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Plugin) OnBuildResource(ctx context.Context, taskCtx core.TaskExecutionContext) *Plugin_BuildResource {
	c := _m.On("BuildResource", ctx, taskCtx)
	return &Plugin_BuildResource{Call: c}
}

func (_m *Plugin) OnBuildResourceMatch(matchers ...interface{}) *Plugin_BuildResource {
	c := _m.On("BuildResource", matchers...)
	return &Plugin_BuildResource{Call: c}
}

// BuildResource provides a mock function with given fields: ctx, taskCtx
func (_m *Plugin) BuildResource(ctx context.Context, taskCtx core.TaskExecutionContext) (client.Object, error) {
	ret := _m.Called(ctx, taskCtx)

	var r0 client.Object
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionContext) client.Object); ok {
		r0 = rf(ctx, taskCtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.TaskExecutionContext) error); ok {
		r1 = rf(ctx, taskCtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Plugin_GetTaskPhase struct {
	*mock.Call
}

func (_m Plugin_GetTaskPhase) Return(_a0 core.PhaseInfo, _a1 error) *Plugin_GetTaskPhase {
	return &Plugin_GetTaskPhase{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Plugin) OnGetTaskPhase(ctx context.Context, resource client.Object) *Plugin_GetTaskPhase {
	c := _m.On("GetTaskPhase", ctx, resource)
	return &Plugin_GetTaskPhase{Call: c}
}

func (_m *Plugin) OnGetTaskPhaseMatch(matchers ...interface{}) *Plugin_GetTaskPhase {
	c := _m.On("GetTaskPhase", matchers...)
	return &Plugin_GetTaskPhase{Call: c}
}

// GetTaskPhase provides a mock function with given fields: ctx, resource
func (_m *Plugin) GetTaskPhase(ctx context.Context, resource client.Object) (core.PhaseInfo, error) {
	ret := _m.Called(ctx, resource)

	var r0 core.PhaseInfo
	if rf, ok := ret.Get(0).(func(context.Context, client.Object) core.PhaseInfo); ok {
		r0 = rf(ctx, resource)
	} else {
		r0 = ret.Get(0).(core.PhaseInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, client.Object) error); ok {
		r1 = rf(ctx, resource)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
