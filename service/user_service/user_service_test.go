package user_service_test

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/user_repository"
	"myGram/service/user_service"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Add_Success(t *testing.T) {
	userPayload := &dto.NewUserRequest{
		Username: "monday",
		Email:    "monday.day@email.com",
		Age:      21,
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Create = func(userPayload *entity.User) (*dto.UserResponse, errs.Error) {
		return &dto.UserResponse{}, nil
	}

	response, err := userService.Add(userPayload)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestUserService_Add_Fail(t *testing.T) {
	userPayload := &dto.NewUserRequest{
		Username: "monday",
		Email:    "monday.day@email.com",
		Age:      21,
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.Create = func(userPayload *entity.User) (*dto.UserResponse, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.Add(userPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestUserService_Add_PayloadNotValid_Fail(t *testing.T) {
	userPayload := &dto.NewUserRequest{
		Username: "monday",
		Age:      21,
		Password: "rahasia",
	}

	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	response, err := userService.Add(userPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestUserService_Get_Success(t *testing.T) {
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

func TestUserService_Get_PayloadNotValid_Fail(t *testing.T) {
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

func TestUserService_Get_EmailNotValid_Fail(t *testing.T) {
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

func TestUserService_Get_ServerError_Fail(t *testing.T) {
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

func TestUserService_Get_PasswordNotValid_Fail(t *testing.T) {
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
			Password: "password rahasia here",
		}

		return user, nil
	}

	response, err := userService.Get(userPayload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_Delete_Success(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return &entity.User{
			Id:        1,
			Username:  "monday",
			Email:     "monday.day@weeekly.com",
			Password:  "rahasia",
			Age:       21,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	user_repository.Delete = func(userId int) errs.Error {
		return nil
	}

	response, err := userService.Remove(1)

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestUserService_DeleteUserNotFound_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return nil, errs.NewNotFoundError("user not found")
	}

	response, err := userService.Remove(1)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_DeleteUserServer_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return nil, errs.NewInternalServerError("somthing went wrong")
	}

	response, err := userService.Remove(1)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_DeleteUserIdNotValid_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return &entity.User{}, nil
	}

	response, err := userService.Remove(1)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_DeleteUserServerError_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return &entity.User{
			Id:        1,
			Username:  "monday",
			Email:     "monday.day@weeekly.com",
			Password:  "encrypt password here",
			Age:       21,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	user_repository.Delete = func(userId int) errs.Error {
		return errs.NewInternalServerError("somthing went wrong")
	}

	response, err := userService.Remove(1)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestUserService_Edit_InvalidBodyRequest_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	userPaylod := &dto.UserUpdateRequest{}

	response, err := userService.Edit(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestUserService_Edit_GetUserNotFound_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	userPaylod := &dto.UserUpdateRequest{
		Username: "lorem",
		Email:    "lorem@email.com",
	}

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return nil, errs.NewNotFoundError("user not found")
	}

	response, err := userService.Edit(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestUserService_Edit_GetUserServerError_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	userPaylod := &dto.UserUpdateRequest{
		Username: "lorem",
		Email:    "lorem@email.com",
	}

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.Edit(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestUserService_Edit_InvalidUser_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	userPaylod := &dto.UserUpdateRequest{
		Username: "lorem",
		Email:    "lorem@email.com",
	}

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return &entity.User{
			Id:        99,
			Username:  "lorem",
			Email:     "lorem@email.com",
			Password:  "secret",
			Age:       22,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	response, err := userService.Edit(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestUserService_Edit_UpdateUserServerError_Fail(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	userPaylod := &dto.UserUpdateRequest{
		Username: "lorem",
		Email:    "lorem@email.com",
	}

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return &entity.User{
			Id:        1,
			Username:  "lorem",
			Email:     "lorem@email.com",
			Password:  "secret",
			Age:       22,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	user_repository.Update = func(userPayload *entity.User) (*dto.UserUpdateResponse, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.Edit(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestUserService_Edit_Success(t *testing.T) {
	userRepo := user_repository.NewUserRepositoryMock()
	userService := user_service.NewUserService(userRepo)

	userPaylod := &dto.UserUpdateRequest{
		Username: "lorem",
		Email:    "lorem@email.com",
	}

	user_repository.FetchById = func(userId int) (*entity.User, errs.Error) {
		return &entity.User{
			Id:        1,
			Username:  "lorem",
			Email:     "lorem@email.com",
			Password:  "secret",
			Age:       22,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	user_repository.Update = func(userPayload *entity.User) (*dto.UserUpdateResponse, errs.Error) {
		return &dto.UserUpdateResponse{}, nil
	}

	response, err := userService.Edit(1, userPaylod)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
