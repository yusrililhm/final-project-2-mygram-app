package user_service

import (
	"myGram/dto"
	"myGram/pkg/errs"
)

type userServiceMock struct {
}

var (
	Add    func(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error)
	Get    func(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, errs.Error)
	Edit   func(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, errs.Error)
	Remove func(userId int) (*dto.GetUserResponse, errs.Error)
)

func NewUserServiceMock() UserService {
	return &userServiceMock{}
}

// Add implements UserService.
func (usm *userServiceMock) Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error) {
	return Add(userPayload)
}

// Get implements UserService.
func (usm *userServiceMock) Get(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, errs.Error) {
	return Get(userPayload)
}

// Edit implements UserService.
func (usm *userServiceMock) Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, errs.Error) {
	return Edit(userId, userPayload)
}

// Remove implements UserService.
func (usm *userServiceMock) Remove(userId int) (*dto.GetUserResponse, errs.Error) {
	return Remove(userId)
}
