package user_service

import (
	"myGram/dto"
	"myGram/pkg/errs"
)

type userServiceMock struct {
}

var (
	Add func(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error)
)

func NewUserServiceMock() UserService {
	return &userServiceMock{}
}

// Add implements UserService.
func (usm *userServiceMock) Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error) {
	return Add(userPayload)
}
