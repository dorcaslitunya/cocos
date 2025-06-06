// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Server is an autogenerated mock type for the Server type
type Server struct {
	mock.Mock
}

type Server_Expecter struct {
	mock *mock.Mock
}

func (_m *Server) EXPECT() *Server_Expecter {
	return &Server_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with no fields
func (_m *Server) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Server_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *Server_Expecter) Start() *Server_Start_Call {
	return &Server_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *Server_Start_Call) Run(run func()) *Server_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Server_Start_Call) Return(_a0 error) *Server_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Server_Start_Call) RunAndReturn(run func() error) *Server_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with no fields
func (_m *Server) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Server_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Server_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Server_Expecter) Stop() *Server_Stop_Call {
	return &Server_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Server_Stop_Call) Run(run func()) *Server_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Server_Stop_Call) Return(_a0 error) *Server_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Server_Stop_Call) RunAndReturn(run func() error) *Server_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewServer creates a new instance of Server. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Server {
	mock := &Server{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
