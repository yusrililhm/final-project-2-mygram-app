package user_service

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/user_repository"
	"net/http"
)

type UserService interface {
	Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error)
}

type userServiceImpl struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

// Add implements UserService.
func (userService *userServiceImpl) Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error) {

	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Age:      userPayload.Age,
		Password: userPayload.Password,
	}

	user.HashPassword()

	id, err := userService.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusCreated,
		Message:    "create new user successfully",
		Data: dto.UserResponse{
			Id:       id,
			Username: user.Username,
			Email:    user.Email,
			Age:      user.Age,
		},
	}, nil
}
