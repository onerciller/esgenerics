// Code generated by mockery v2.36.0. DO NOT EDIT.

package esgenerics

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockQueryHandler is an autogenerated mock type for the QueryHandler type
type MockQueryHandler struct {
	mock.Mock
}

// BuildQuery provides a mock function with given fields: ctx
func (_m *MockQueryHandler) BuildQuery(ctx context.Context) QueryMap {
	ret := _m.Called(ctx)

	var r0 QueryMap
	if rf, ok := ret.Get(0).(func(context.Context) QueryMap); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(QueryMap)
		}
	}

	return r0
}

// NewMockQueryHandler creates a new instance of MockQueryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueryHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueryHandler {
	mock := &MockQueryHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
