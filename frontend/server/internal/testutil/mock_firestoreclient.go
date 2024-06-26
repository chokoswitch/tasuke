// Code generated by mockery v2.43.0. DO NOT EDIT.

package testutil

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockFirestoreClient is an autogenerated mock type for the Client type
type MockFirestoreClient struct {
	mock.Mock
}

type MockFirestoreClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFirestoreClient) EXPECT() *MockFirestoreClient_Expecter {
	return &MockFirestoreClient_Expecter{mock: &_m.Mock}
}

// SetDocument provides a mock function with given fields: ctx, collection, path, data
func (_m *MockFirestoreClient) SetDocument(ctx context.Context, collection string, path string, data interface{}) error {
	ret := _m.Called(ctx, collection, path, data)

	if len(ret) == 0 {
		panic("no return value specified for SetDocument")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, collection, path, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFirestoreClient_SetDocument_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDocument'
type MockFirestoreClient_SetDocument_Call struct {
	*mock.Call
}

// SetDocument is a helper method to define mock.On call
//   - ctx context.Context
//   - collection string
//   - path string
//   - data interface{}
func (_e *MockFirestoreClient_Expecter) SetDocument(ctx interface{}, collection interface{}, path interface{}, data interface{}) *MockFirestoreClient_SetDocument_Call {
	return &MockFirestoreClient_SetDocument_Call{Call: _e.mock.On("SetDocument", ctx, collection, path, data)}
}

func (_c *MockFirestoreClient_SetDocument_Call) Run(run func(ctx context.Context, collection string, path string, data interface{})) *MockFirestoreClient_SetDocument_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(interface{}))
	})
	return _c
}

func (_c *MockFirestoreClient_SetDocument_Call) Return(_a0 error) *MockFirestoreClient_SetDocument_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFirestoreClient_SetDocument_Call) RunAndReturn(run func(context.Context, string, string, interface{}) error) *MockFirestoreClient_SetDocument_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFirestoreClient creates a new instance of MockFirestoreClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFirestoreClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFirestoreClient {
	mock := &MockFirestoreClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
