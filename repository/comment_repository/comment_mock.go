package comment_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type commentRepositoryMock struct {
}

var (
	AddComment func(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error)
)

func NewCommentRepositoryMock() CommentRepository {
	return &commentRepositoryMock{}
}

// AddComment implements CommentRepository.
func (crm *commentRepositoryMock) AddComment(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error) {
	return AddComment(commentPayload)
}
