// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "prototype2/domain"

	mock "github.com/stretchr/testify/mock"
)

// PostRepository is an autogenerated mock type for the PostRepository type
type PostRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: post
func (_m *PostRepository) Delete(post *domain.Post) error {
	ret := _m.Called(post)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Post) error); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *PostRepository) FindAll() ([]domain.Post, error) {
	ret := _m.Called()

	var r0 []domain.Post
	if rf, ok := ret.Get(0).(func() []domain.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Post)
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

// Migrate provides a mock function with given fields:
func (_m *PostRepository) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: post
func (_m *PostRepository) Save(post *domain.Post) (*domain.Post, error) {
	ret := _m.Called(post)

	var r0 *domain.Post
	if rf, ok := ret.Get(0).(func(*domain.Post) *domain.Post); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// func (mock *MockRepository) Save(post *domain.Post) (*domain.Post, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.(*domain.Post), args.Error(1)
// }
