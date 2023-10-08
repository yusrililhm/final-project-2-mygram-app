package comment_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type commentRepositoryMock struct {
}

var (
	AddComment     func(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error)
	GetComments    func() ([]CommentUserPhotoMapped, errs.Error)
	GetCommentById func(commentId int) (*CommentUserPhotoMapped, errs.Error)
)

func NewCommentRepositoryMock() CommentRepository {
	return &commentRepositoryMock{}
}

// AddComment implements CommentRepository.
func (crm *commentRepositoryMock) AddComment(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error) {
	return AddComment(commentPayload)
}

// GetComments implements CommentRepository.
func (crm *commentRepositoryMock) GetComments() ([]CommentUserPhotoMapped, errs.Error) {
	return GetComments()
}

// GetCommentById implements CommentRepository.
func (crm *commentRepositoryMock) GetCommentById(commentId int) (*CommentUserPhotoMapped, errs.Error) {
	return GetCommentById(commentId)
}
