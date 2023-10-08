package handler

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/service/comment_service"

	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	AddComment(ctx *gin.Context)
	GetComments(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHandlerImpl struct {
	commentService comment_service.CommentService
}

func NewCommentHandler(commentService comment_service.CommentService) CommentHandler {
	return &commentHandlerImpl{
		commentService: commentService,
	}
}

// AddComment implements CommentHandler.
func (commentHandler *commentHandlerImpl) AddComment(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)
	commentPayload := &dto.NewCommentRequest{}

	if err := ctx.ShouldBindJSON(commentPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := commentHandler.commentService.AddComment(user.Id, commentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteComment implements CommentHandler.
func (commentHandler *commentHandlerImpl) DeleteComment(ctx *gin.Context) {
	panic("unimplemented")
}

// GetComments implements CommentHandler.
func (commentHandler *commentHandlerImpl) GetComments(ctx *gin.Context) {
	response, err := commentHandler.commentService.GetComments()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateComment implements CommentHandler.
func (commentHandler *commentHandlerImpl) UpdateComment(ctx *gin.Context) {
	panic("unimplemented")
}
