package photo_service_test

import (
	"net/http"
	"testing"

	"myGram/dto"
	"myGram/entity"

	"myGram/pkg/errs"

	"myGram/repository/photo_repository"
	"myGram/service/photo_service"

	"github.com/stretchr/testify/assert"
)

func TestPhotoService_AddPhoto_BodyNotValid_Fail(t *testing.T) {

	photoPayload := &dto.NewPhotoRequest{
		Title:   "daily monday",
		Caption: "i'm monday",
	}

	photoRepository := photo_repository.NewPhotoRepositoryMock()
	photoService := photo_service.NewPhotoService(photoRepository)

	response, err := photoService.AddPhoto(1, photoPayload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestPhotoService_AddPhoto_ServerError_Fail(t *testing.T) {

	photoPayload := &dto.NewPhotoRequest{
		Title:    "daily monday",
		PhotoUrl: "https://google.com/monday",
		Caption:  "i'm monday",
	}

	photoRepository := photo_repository.NewPhotoRepositoryMock()
	photoService := photo_service.NewPhotoService(photoRepository)

	photo_repository.AddPhoto = func(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := photoService.AddPhoto(1, photoPayload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Message(), "something went wrong")
}

func TestPhotoService_AddPhoto_Success(t *testing.T) {

	photoPayload := &dto.NewPhotoRequest{
		Title:    "daily monday",
		PhotoUrl: "https://google.com/monday",
		Caption:  "i'm monday",
	}

	photoRepository := photo_repository.NewPhotoRepositoryMock()
	photoService := photo_service.NewPhotoService(photoRepository)

	photo_repository.AddPhoto = func(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error) {
		return &dto.PhotoResponse{}, nil
	}

	response, err := photoService.AddPhoto(1, photoPayload)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestPhotoService_GetPhotos_ServerError_Fail(t *testing.T)  {
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	photoService := photo_service.NewPhotoService(photoRepo)

	photo_repository.GetPhotos = func() ([]photo_repository.PhotoUserMapped, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := photoService.GetPhotos()

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestPhotoService_GetPhotos_PhotoNotFound_Fail(t *testing.T)  {
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	photoService := photo_service.NewPhotoService(photoRepo)

	photo_repository.GetPhotos = func() ([]photo_repository.PhotoUserMapped, errs.Error) {
		return nil, errs.NewNotFoundError("photos not found")
	}

	response, err := photoService.GetPhotos()

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestPhotoService_GetPhotos_Success(t *testing.T)  {
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	photoService := photo_service.NewPhotoService(photoRepo)

	photo_repository.GetPhotos = func() ([]photo_repository.PhotoUserMapped, errs.Error) {
		return []photo_repository.PhotoUserMapped{}, nil
	}

	response, err := photoService.GetPhotos()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
