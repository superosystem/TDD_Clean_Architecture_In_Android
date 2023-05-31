// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	users "github.com/superosystem/bantumanten-backend/src/businesses/users"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: domain
func (_m *Repository) Create(domain *users.Domain) error {
	ret := _m.Called(domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(*users.Domain) error); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ID
func (_m *Repository) Delete(ID string) bool {
	ret := _m.Called(ID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() []users.Domain {
	ret := _m.Called()

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func() []users.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
		}
	}

	return r0
}

// GetByEmail provides a mock function with given fields: email
func (_m *Repository) GetByEmail(email string) (*users.Domain, error) {
	ret := _m.Called(email)

	var r0 *users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*users.Domain, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *users.Domain); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ID
func (_m *Repository) GetByID(ID string) (*users.Domain, error) {
	ret := _m.Called(ID)

	var r0 *users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*users.Domain, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(string) *users.Domain); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsExistUser provides a mock function with given fields: ID, email
func (_m *Repository) IsExistUser(ID string, email string) bool {
	ret := _m.Called(ID, email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(ID, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Update provides a mock function with given fields: ID, domain
func (_m *Repository) Update(ID string, domain *users.Domain) (*users.Domain, error) {
	ret := _m.Called(ID, domain)

	var r0 *users.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *users.Domain) (*users.Domain, error)); ok {
		return rf(ID, domain)
	}
	if rf, ok := ret.Get(0).(func(string, *users.Domain) *users.Domain); ok {
		r0 = rf(ID, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *users.Domain) error); ok {
		r1 = rf(ID, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
