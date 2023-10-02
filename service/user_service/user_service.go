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
	Get(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, errs.Error)
	Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, errs.Error)
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

// Get implements UserService.
func (us *userServiceImpl) Get(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, errs.Error) {

	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user, err := us.userRepo.Fetch(userPayload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(userPayload.Password)

	if isValidPassword == false {
		return nil, errs.NewBadRequestError("invalid email/password")
	}

	token := user.GenerateToken()

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "successfully loged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}, nil
}

// Edit implements UserService.
func (userService *userServiceImpl) Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, errs.Error) {

	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user, err := userService.userRepo.FetchById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("invalid email/password")
		}
		return nil, err
	}

	if user.Id != userId {
		return nil, errs.NewUnauthenticatedError("invalid user")
	}

	usr := &entity.User{
		Id:       userId,
		Email:    userPayload.Email,
		Username: userPayload.Username,
	}
	err = userService.userRepo.Update(usr)

	if err != nil {
		return nil, err
	}

	user, err = userService.userRepo.FetchById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("invalid email/password")
		}
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "user updated successfully",
		Data: dto.UserUpdateResponse{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			Age:       user.Age,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}
