package comment_service

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/comment_repository"
	"myGram/repository/photo_repository"
	"net/http"
)

type CommentService interface {
	AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error)
	GetComments() (*dto.GetCommentResponse, errs.Error)
}

type commentServiceImpl struct {
	photoRepo   photo_repository.PhotoRepository
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository, photoRepo photo_repository.PhotoRepository) CommentService {
	return &commentServiceImpl{
		photoRepo:   photoRepo,
		commentRepo: commentRepo,
	}
}

// AddComment implements CommentService.
func (commentService *commentServiceImpl) AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error) {

	err := helper.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	_, err = commentService.photoRepo.GetPhotoId(commentPayload.PhotoId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	comment := &entity.Comment{
		UserId:  userId,
		PhotoId: commentPayload.PhotoId,
		Message: commentPayload.Message,
	}

	response, err := commentService.commentRepo.AddComment(comment)

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusCreated,
		Message:    "new comment successfully added",
		Data:       response,
	}, nil
}

// GetComments implements CommentService.
func (commentService *commentServiceImpl) GetComments() (*dto.GetCommentResponse, errs.Error) {

	data, err := commentService.commentRepo.GetComments()

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "fetch comments successfully",
		Data:       data,
	}, nil
}
