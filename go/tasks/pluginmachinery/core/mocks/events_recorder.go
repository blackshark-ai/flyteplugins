// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import "context"
import "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
import "github.com/stretchr/testify/mock"

// EventsRecorder is an autogenerated mock type for the EventsRecorder type
type EventsRecorder struct {
	mock.Mock
}

type EventsRecorder_RecordRaw struct {
	*mock.Call
}

func (_m EventsRecorder_RecordRaw) Return(_a0 error) *EventsRecorder_RecordRaw {
	return &EventsRecorder_RecordRaw{Call: _m.Call.Return(_a0)}
}

func (_m *EventsRecorder) OnRecordRaw(ctx context.Context, ev core.PhaseInfo) *EventsRecorder_RecordRaw {
	c := _m.On("RecordRaw")
	return &EventsRecorder_RecordRaw{Call: c}
}
func (_m *EventsRecorder) OnRecordRawMatch(matchers ...interface{}) *EventsRecorder_RecordRaw {
	c := _m.On("RecordRaw", matchers...)
	return &EventsRecorder_RecordRaw{Call: c}
}

// RecordRaw provides a mock function with given fields: ctx, ev
func (_m *EventsRecorder) RecordRaw(ctx context.Context, ev core.PhaseInfo) error {
	ret := _m.Called(ctx, ev)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.PhaseInfo) error); ok {
		r0 = rf(ctx, ev)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
