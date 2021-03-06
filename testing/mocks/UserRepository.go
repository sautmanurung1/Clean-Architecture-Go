// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	users "Belajar-Go-Echo/domains/users"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserRepository) CreateUser(user users.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUser provides a mock function with given fields:
func (_m *UserRepository) GetAllUser() []users.User {
	ret := _m.Called()

	var r0 []users.User
	if rf, ok := ret.Get(0).(func() []users.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.User)
		}
	}

	return r0
}

// GetUserById provides a mock function with given fields: id
func (_m *UserRepository) GetUserById(id int) (users.User, error) {
	ret := _m.Called(id)

	var r0 users.User
	if rf, ok := ret.Get(0).(func(int) users.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
