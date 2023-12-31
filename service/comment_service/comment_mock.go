package comment_service

import (
	"myGram/dto"
	"myGram/pkg/errs"
)

type commentServiceMock struct {
}

var (
	AddComment    func(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error)
	GetComments   func() (*dto.GetCommentResponse, errs.Error)
	DeleteComment func(commentId int) (*dto.GetCommentResponse, errs.Error)
	UpdateComment func(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.GetCommentResponse, errs.Error)
)

func NewCommentServiceMock() CommentService {
	return &commentServiceMock{}
}

// AddComment implements CommentService.
func (csm *commentServiceMock) AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error) {
	return AddComment(userId, commentPayload)
}

// GetComments implements CommentService.
func (csm *commentServiceMock) GetComments() (*dto.GetCommentResponse, errs.Error) {
	return GetComments()
}

// DeleteComment implements CommentService.
func (csm *commentServiceMock) DeleteComment(commentId int) (*dto.GetCommentResponse, errs.Error) {
	return DeleteComment(commentId)
}

// UpdateComment implements CommentService.
func (csm *commentServiceMock) UpdateComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.GetCommentResponse, errs.Error) {
	return UpdateComment(commentId, commentPayload)
}
