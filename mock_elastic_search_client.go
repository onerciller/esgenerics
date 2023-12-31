// Code generated by mockery v2.36.0. DO NOT EDIT.

package esgenerics

import (
	context "context"

	esapi "github.com/elastic/go-elasticsearch/v7/esapi"
	mock "github.com/stretchr/testify/mock"
)

// MockElasticSearchClient is an autogenerated mock type for the ElasticSearchClient type
type MockElasticSearchClient struct {
	mock.Mock
}

// Search provides a mock function with given fields: ctx, indexName, queryMap
func (_m *MockElasticSearchClient) Search(ctx context.Context, indexName string, queryMap QueryMap) (*esapi.Response, error) {
	ret := _m.Called(ctx, indexName, queryMap)

	var r0 *esapi.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, QueryMap) (*esapi.Response, error)); ok {
		return rf(ctx, indexName, queryMap)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, QueryMap) *esapi.Response); ok {
		r0 = rf(ctx, indexName, queryMap)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*esapi.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, QueryMap) error); ok {
		r1 = rf(ctx, indexName, queryMap)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockElasticSearchClient creates a new instance of MockElasticSearchClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockElasticSearchClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockElasticSearchClient {
	mock := &MockElasticSearchClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
