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
		Email: "monday.day@email.com",
		Age: 21,
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Create = func(userPayload *entity.User) (int, errs.Error) {
		return 1, nil
	}

	response, err := userService.Add(userPayload)
	expected := &dto.GetUserResponse{
		StatusCode: http.StatusCreated,
		Message: "create new user successfully",
		Data: dto.UserResponse{
			Id: 1,
			Username: "monday",
			Email: "monday.day@email.com",
			Age: 21,
		},
	}

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, expected, response)
}

func TestUserService_AddFail(t *testing.T) {
	userPayload := &dto.NewUserRequest{}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Create = func(userPayload *entity.User) (int, errs.Error) {
		return 0, errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.Add(userPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
}
