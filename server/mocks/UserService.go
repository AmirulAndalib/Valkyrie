// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/sentrionic/valkyrie/model"
	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	testing "testing"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// ChangeAvatar provides a mock function with given fields: header, directory
func (_m *UserService) ChangeAvatar(header *multipart.FileHeader, directory string) (string, error) {
	ret := _m.Called(header, directory)

	var r0 string
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, string) string); ok {
		r0 = rf(header, directory)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*multipart.FileHeader, string) error); ok {
		r1 = rf(header, directory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangePassword provides a mock function with given fields: currentPassword, newPassword, user
func (_m *UserService) ChangePassword(currentPassword string, newPassword string, user *model.User) error {
	ret := _m.Called(currentPassword, newPassword, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *model.User) error); ok {
		r0 = rf(currentPassword, newPassword, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteImage provides a mock function with given fields: key
func (_m *UserService) DeleteImage(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForgotPassword provides a mock function with given fields: ctx, user
func (_m *UserService) ForgotPassword(ctx context.Context, user *model.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *UserService) Get(id string) (*model.User, error) {
	ret := _m.Called(id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserService) GetByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFriendAndGuildIds provides a mock function with given fields: userId
func (_m *UserService) GetFriendAndGuildIds(userId string) (*[]string, error) {
	ret := _m.Called(userId)

	var r0 *[]string
	if rf, ok := ret.Get(0).(func(string) *[]string); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRequestCount provides a mock function with given fields: userId
func (_m *UserService) GetRequestCount(userId string) (*int64, error) {
	ret := _m.Called(userId)

	var r0 *int64
	if rf, ok := ret.Get(0).(func(string) *int64); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEmailAlreadyInUse provides a mock function with given fields: email
func (_m *UserService) IsEmailAlreadyInUse(email string) bool {
	ret := _m.Called(email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *UserService) Login(email string, password string) (*model.User, error) {
	ret := _m.Called(email, password)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string, string) *model.User); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: user
func (_m *UserService) Register(user *model.User) (*model.User, error) {
	ret := _m.Called(user)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: ctx, password, token
func (_m *UserService) ResetPassword(ctx context.Context, password string, token string) (*model.User, error) {
	ret := _m.Called(ctx, password, token)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.User); ok {
		r0 = rf(ctx, password, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, password, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccount provides a mock function with given fields: user
func (_m *UserService) UpdateAccount(user *model.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserService creates a new instance of UserService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t testing.TB) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
