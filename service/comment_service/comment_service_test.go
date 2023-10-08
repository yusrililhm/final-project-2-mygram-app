package comment_service_test

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/comment_repository"
	"myGram/repository/photo_repository"
	"myGram/service/comment_service"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentService_AddComment_InvalidBodyRequest_Fail(t *testing.T) {

	commentPayload := &dto.NewCommentRequest{}

	commentRepo := comment_repository.NewCommentRepositoryMock()
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)

	response, err := commentService.AddComment(1, commentPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestCommentService_AddComment_PhotoNotFound_Fail(t *testing.T) {

	commentPayload := &dto.NewCommentRequest{
		PhotoId: 99,
		Message: "lorem ipsum",
	}

	commentRepo := comment_repository.NewCommentRepositoryMock()
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)

	photo_repository.GetPhotoId = func(photoId int) (*photo_repository.PhotoUserMapped, errs.Error) {
		return nil, errs.NewNotFoundError("photo not found")
	}

	response, err := commentService.AddComment(1, commentPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestCommentService_AddComment_GetPhotoError_Fail(t *testing.T) {

	commentPayload := &dto.NewCommentRequest{
		PhotoId: 99,
		Message: "lorem ipsum",
	}

	commentRepo := comment_repository.NewCommentRepositoryMock()
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)

	photo_repository.GetPhotoId = func(photoId int) (*photo_repository.PhotoUserMapped, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := commentService.AddComment(1, commentPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestCommentService_AddComment_ServerError_Fail(t *testing.T) {

	commentPayload := &dto.NewCommentRequest{
		PhotoId: 99,
		Message: "lorem ipsum",
	}

	commentRepo := comment_repository.NewCommentRepositoryMock()
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)

	photo_repository.GetPhotoId = func(photoId int) (*photo_repository.PhotoUserMapped, errs.Error) {
		return &photo_repository.PhotoUserMapped{}, nil
	}

	comment_repository.AddComment = func(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := commentService.AddComment(1, commentPayload)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestCommentService_AddComment_Success(t *testing.T) {

	commentPayload := &dto.NewCommentRequest{
		PhotoId: 99,
		Message: "lorem ipsum",
	}

	commentRepo := comment_repository.NewCommentRepositoryMock()
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)

	photo_repository.GetPhotoId = func(photoId int) (*photo_repository.PhotoUserMapped, errs.Error) {
		return &photo_repository.PhotoUserMapped{}, nil
	}

	comment_repository.AddComment = func(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error) {
		return &dto.NewCommentResponse{}, nil
	}

	response, err := commentService.AddComment(1, commentPayload)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestCommentService_GetComments_ServerError_Fail(t *testing.T) {
	commentRepo := comment_repository.NewCommentRepositoryMock()
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)

	comment_repository.GetComments = func() ([]comment_repository.CommentUserPhotoMapped, errs.Error) {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response, err := commentService.GetComments()

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestCommentService_GetComments_Success(t *testing.T) {
	commentRepo := comment_repository.NewCommentRepositoryMock()
	photoRepo := photo_repository.NewPhotoRepositoryMock()
	commentService := comment_service.NewCommentService(commentRepo, photoRepo)

	comment_repository.GetComments = func() ([]comment_repository.CommentUserPhotoMapped, errs.Error) {
		return []comment_repository.CommentUserPhotoMapped{}, nil
	}

	response, err := commentService.GetComments()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
