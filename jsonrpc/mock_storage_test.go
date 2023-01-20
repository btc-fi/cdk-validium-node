// Code generated by mockery v2.16.0. DO NOT EDIT.

package jsonrpc

import (
	websocket "github.com/gorilla/websocket"
	mock "github.com/stretchr/testify/mock"
)

// storageMock is an autogenerated mock type for the storageInterface type
type storageMock struct {
	mock.Mock
}

// GetAllBlockFiltersWithWSConn provides a mock function with given fields:
func (_m *storageMock) GetAllBlockFiltersWithWSConn() ([]*Filter, error) {
	ret := _m.Called()

	var r0 []*Filter
	if rf, ok := ret.Get(0).(func() []*Filter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Filter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllLogFiltersWithWSConn provides a mock function with given fields:
func (_m *storageMock) GetAllLogFiltersWithWSConn() ([]*Filter, error) {
	ret := _m.Called()

	var r0 []*Filter
	if rf, ok := ret.Get(0).(func() []*Filter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Filter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFilter provides a mock function with given fields: filterID
func (_m *storageMock) GetFilter(filterID string) (*Filter, error) {
	ret := _m.Called(filterID)

	var r0 *Filter
	if rf, ok := ret.Get(0).(func(string) *Filter); ok {
		r0 = rf(filterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Filter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlockFilter provides a mock function with given fields: wsConn
func (_m *storageMock) NewBlockFilter(wsConn *websocket.Conn) (string, error) {
	ret := _m.Called(wsConn)

	var r0 string
	if rf, ok := ret.Get(0).(func(*websocket.Conn) string); ok {
		r0 = rf(wsConn)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*websocket.Conn) error); ok {
		r1 = rf(wsConn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLogFilter provides a mock function with given fields: wsConn, filter
func (_m *storageMock) NewLogFilter(wsConn *websocket.Conn, filter LogFilter) (string, error) {
	ret := _m.Called(wsConn, filter)

	var r0 string
	if rf, ok := ret.Get(0).(func(*websocket.Conn, LogFilter) string); ok {
		r0 = rf(wsConn, filter)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*websocket.Conn, LogFilter) error); ok {
		r1 = rf(wsConn, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPendingTransactionFilter provides a mock function with given fields: wsConn
func (_m *storageMock) NewPendingTransactionFilter(wsConn *websocket.Conn) (string, error) {
	ret := _m.Called(wsConn)

	var r0 string
	if rf, ok := ret.Get(0).(func(*websocket.Conn) string); ok {
		r0 = rf(wsConn)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*websocket.Conn) error); ok {
		r1 = rf(wsConn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UninstallFilter provides a mock function with given fields: filterID
func (_m *storageMock) UninstallFilter(filterID string) error {
	ret := _m.Called(filterID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filterID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UninstallFilterByWSConn provides a mock function with given fields: wsConn
func (_m *storageMock) UninstallFilterByWSConn(wsConn *websocket.Conn) error {
	ret := _m.Called(wsConn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*websocket.Conn) error); ok {
		r0 = rf(wsConn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFilterLastPoll provides a mock function with given fields: filterID
func (_m *storageMock) UpdateFilterLastPoll(filterID string) error {
	ret := _m.Called(filterID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filterID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewStorageMock interface {
	mock.TestingT
	Cleanup(func())
}

// newStorageMock creates a new instance of storageMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newStorageMock(t mockConstructorTestingTnewStorageMock) *storageMock {
	mock := &storageMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
