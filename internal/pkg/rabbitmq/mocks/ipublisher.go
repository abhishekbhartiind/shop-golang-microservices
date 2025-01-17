// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IPublisher is an autogenerated mock type for the IPublisher type
type IPublisher struct {
	mock.Mock
}

// IsPublished provides a mock function with given fields: msg
func (_m *IPublisher) IsPublished(msg interface{}) bool {
	ret := _m.Called(msg)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PublishMessage provides a mock function with given fields: ctx, msg
func (_m *IPublisher) PublishMessage(ctx context.Context, msg interface{}) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPublisher creates a new instance of IPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPublisher(t mockConstructorTestingTNewIPublisher) *IPublisher {
	mock := &IPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
