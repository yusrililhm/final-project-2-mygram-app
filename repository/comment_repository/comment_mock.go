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
	DeleteComment  func(commentId int) errs.Error
	UpdateComment         func(commentId int, commentPayload *entity.Comment) (*dto.PhotoUpdateResponse, errs.Error)
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

// DeleteComment implements CommentRepository.
func (crm *commentRepositoryMock) DeleteComment(commentId int) errs.Error {
	return DeleteComment(commentId)
}

// Update implements CommentRepository.
func (crm *commentRepositoryMock) UpdateComment(commentId int, commentPayload *entity.Comment) (*dto.PhotoUpdateResponse, errs.Error) {
	return UpdateComment(commentId, commentPayload)
}
