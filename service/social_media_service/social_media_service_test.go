package social_media_service_test

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/social_media_repository"
	"myGram/service/social_media_service"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSocialMediaService_Add_InvalidRequest_Fail(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	userPaylod := &dto.NewSocialMediaRequest{}
	response, err := socialMediaService.AddSocialMedia(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestSocialMediaService_Add_ServerError_Fail(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	userPaylod := &dto.NewSocialMediaRequest{
		Name:           "lorem",
		SocialMediaUrl: "lorem ipsum",
	}

	social_media_repository.AddSocialMedia = func(socialMediaPayload *entity.SocialMedia) (*dto.NewSocialMediaResponse, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	response, err := socialMediaService.AddSocialMedia(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSocialMediaService_Add_Success(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	userPaylod := &dto.NewSocialMediaRequest{
		Name:           "lorem",
		SocialMediaUrl: "lorem ipsum",
	}

	social_media_repository.AddSocialMedia = func(socialMediaPayload *entity.SocialMedia) (*dto.NewSocialMediaResponse, errs.Error) {
		return &dto.NewSocialMediaResponse{}, nil
	}
	response, err := socialMediaService.AddSocialMedia(1, userPaylod)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestSocialMediaService_Get_ServerError_Fail(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	social_media_repository.GetSocialMedias = func() ([]*dto.GetSocialMedia, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := socialMediaService.GetSocialMedias()

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSocialMediaService_Get_Success(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	social_media_repository.GetSocialMedias = func() ([]*dto.GetSocialMedia, errs.Error) {
		return []*dto.GetSocialMedia{}, nil
	}

	response, err := socialMediaService.GetSocialMedias()

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestSocialMediaService_Delete_ServerError_Fail(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	social_media_repository.DeleteSocialMedia = func(socialMediaId int) errs.Error {
		return errs.NewInternalServerError("something went wrong")
	}

	response, err := socialMediaService.DeleteSocialMedia(1)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSocialMediaService_Delete_Success(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	social_media_repository.DeleteSocialMedia = func(socialMediaId int) errs.Error {
		return nil
	}

	response, err := socialMediaService.DeleteSocialMedia(1)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestSocialMediaService_Update_InvalidRequest_Fail(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	userPaylod := &dto.UpdateSocialMediaRequest{}
	response, err := socialMediaService.UpdateSocialMedia(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestSocialMediaService_Update_ServerError_Fail(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	userPaylod := &dto.UpdateSocialMediaRequest{
		Name:           "lorem",
		SocialMediaUrl: "lorem ipsum",
	}

	social_media_repository.UpdateSocialMedia = func(socialMediaId int, socialMediaPayload *entity.SocialMedia) (*dto.SocialMediaUpdateResponse, errs.Error) {
		return nil, errs.NewInternalServerError("somthing went wrong")
	}

	response, err := socialMediaService.UpdateSocialMedia(1, userPaylod)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSocialMediaService_Update_Success(t *testing.T) {
	socialMediaRepo := social_media_repository.NewSocialMediaMock()
	socialMediaService := social_media_service.NewSocialMediaService(socialMediaRepo)

	userPaylod := &dto.UpdateSocialMediaRequest{
		Name:           "lorem",
		SocialMediaUrl: "lorem ipsum",
	}

	social_media_repository.UpdateSocialMedia = func(socialMediaId int, socialMediaPayload *entity.SocialMedia) (*dto.SocialMediaUpdateResponse, errs.Error) {
		return &dto.SocialMediaUpdateResponse{}, nil
	}

	response, err := socialMediaService.UpdateSocialMedia(1, userPaylod)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
