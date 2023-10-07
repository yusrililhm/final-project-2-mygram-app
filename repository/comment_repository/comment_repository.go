package comment_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type CommentRepository interface {
	AddComment(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error)
}
