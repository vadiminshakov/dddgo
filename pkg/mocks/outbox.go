// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Outbox is an autogenerated mock type for the Outbox type
type Outbox struct {
	mock.Mock
}

// Save provides a mock function with given fields: key, value
func (_m *Outbox) Save(key string, value []byte) error {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOutbox creates a new instance of Outbox. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOutbox(t interface {
	mock.TestingT
	Cleanup(func())
}) *Outbox {
	mock := &Outbox{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}