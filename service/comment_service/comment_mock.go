package comment_service

import (
	"myGram/dto"
	"myGram/pkg/errs"
)

type commentServiceMock struct {
}

var (
	AddComment func(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error)
)

func NewCommentServiceMock() CommentService {
	return &commentServiceMock{}
}

// AddComment implements CommentService.
func (csm *commentServiceMock) AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error) {
	return AddComment(userId, commentPayload)
}
