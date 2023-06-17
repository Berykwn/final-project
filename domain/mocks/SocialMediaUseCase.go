// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "mygram-api/domain"

	mock "github.com/stretchr/testify/mock"
)

// SocialMediaUseCase is an autogenerated mock type for the SocialMediaUseCase type
type SocialMediaUseCase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *SocialMediaUseCase) Delete(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: _a0, _a1, _a2
func (_m *SocialMediaUseCase) Fetch(_a0 context.Context, _a1 *[]domain.SocialMedia, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *[]domain.SocialMedia, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *SocialMediaUseCase) GetByID(_a0 context.Context, _a1 *domain.SocialMedia, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.SocialMedia, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *SocialMediaUseCase) Store(_a0 context.Context, _a1 *domain.SocialMedia) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.SocialMedia) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *SocialMediaUseCase) Update(_a0 context.Context, _a1 domain.SocialMedia, _a2 string) (domain.SocialMedia, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 domain.SocialMedia
	if rf, ok := ret.Get(0).(func(context.Context, domain.SocialMedia, string) domain.SocialMedia); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(domain.SocialMedia)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.SocialMedia, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSocialMediaUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewSocialMediaUseCase creates a new instance of SocialMediaUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSocialMediaUseCase(t mockConstructorTestingTNewSocialMediaUseCase) *SocialMediaUseCase {
	mock := &SocialMediaUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}