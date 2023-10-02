package user_service_test

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/user_repository"
	"myGram/service/user_service"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_AddSuccess(t *testing.T) {
	userPayload := &dto.NewUserRequest{
		Username: "monday",
		Email:    "monday.day@email.com",
		Age:      21,
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Create = func(userPayload *entity.User) (int, errs.Error) {
		return 1, nil
	}

	response, err := userService.Add(userPayload)
	expected := &dto.GetUserResponse{
		StatusCode: http.StatusCreated,
		Message:    "create new user successfully",
		Data: dto.UserResponse{
			Id:       1,
			Username: "monday",
			Email:    "monday.day@email.com",
			Age:      21,
		},
	}

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, expected, response)
}

func TestUserService_AddFail(t *testing.T) {
	userPayload := &dto.NewUserRequest{
		Username: "monday",
		Email:    "monday.day@email.com",
		Age:      21,
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Create = func(userPayload *entity.User) (int, errs.Error) {
		return 0, errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.Add(userPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestUserService_AddFailPayloadNotValid(t *testing.T) {
	userPayload := &dto.NewUserRequest{
		Username: "monday",
		Age:      21,
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Create = func(userPayload *entity.User) (int, errs.Error) {
		return 0, errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.Add(userPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestUserService_GetSuccess(t *testing.T) {
	userPayload := &dto.UserLoginRequest{
		Email:    "monday.day@email.com",
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Fetch = func(email string) (*entity.User, errs.Error) {
		return &entity.User{
			Id:       6,
			Username: "monday",
			Email:    "monday.day@email.com",
			Age:      21,
			Password: "$2a$10$r7EA6IIKVoh7Pr2KQIN7NeHr/IDHWldIudGdRVeOmBW0wLXte9aqG",
		}, nil
	}

	response, err := userService.Get(userPayload)

	assert.NotNil(t, response)
	assert.Nil(t, err)
}

func TestUserService_GetFailPayloadNotValid(t *testing.T) {
	userPayload := &dto.UserLoginRequest{}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Fetch = func(email string) (*entity.User, errs.Error) {
		return nil, errs.NewBadRequestError("payload not valid")
	}

	response, err := userService.Get(userPayload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_GetFailEmailNotValid(t *testing.T) {
	userPayload := &dto.UserLoginRequest{
		Email:    "monday.dayy@email.com",
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Fetch = func(email string) (*entity.User, errs.Error) {
		return nil, errs.NewNotFoundError("user not found")
	}

	response, err := userService.Get(userPayload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_GetFailRepoError(t *testing.T) {
	userPayload := &dto.UserLoginRequest{
		Email:    "monday.day@email.com",
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Fetch = func(email string) (*entity.User, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.Get(userPayload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_GetFailPasswordNotValid(t *testing.T) {
	userPayload := &dto.UserLoginRequest{
		Email:    "monday.dayy@email.com",
		Password: "rahasiashhhh",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Fetch = func(email string) (*entity.User, errs.Error) {

		user := &entity.User{
			Id:       6,
			Username: "monday",
			Email:    "monday.day@email.com",
			Age:      21,
			Password: "$2a$10$r7EA6IIKVoh7Pr2KQIN7NeHr/IDHWldIudGdRVeOmBW0wLXte9aqG",
		}

		isValidPassword := user.ComparePassword(userPayload.Password)

		if isValidPassword == false {
			return nil, errs.NewBadRequestError("invalid email/password")
		}

		return user, nil
	}

	response, err := userService.Get(userPayload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}
