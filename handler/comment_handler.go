package handler

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/service/comment_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	AddComment(ctx *gin.Context)
	GetComments(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHandlerImpl struct {
	cs comment_service.CommentService
}

func NewCommentHandler(commentService comment_service.CommentService) CommentHandler {
	return &commentHandlerImpl{
		cs: commentService,
	}
}

// AddComment implements CommentHandler.
func (c *commentHandlerImpl) AddComment(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)
	commentPayload := &dto.NewCommentRequest{}

	if err := ctx.ShouldBindJSON(commentPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := c.cs.AddComment(user.Id, commentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteComment implements CommentHandler.
func (c *commentHandlerImpl) DeleteComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	response, err := c.cs.DeleteComment(commentId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetComments implements CommentHandler.
func (c *commentHandlerImpl) GetComments(ctx *gin.Context) {
	response, err := c.cs.GetComments()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateComment implements CommentHandler.
func (c *commentHandlerImpl) UpdateComment(ctx *gin.Context) {
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	commentPayload := &dto.UpdateCommentRequest{}

	if err := ctx.ShouldBindJSON(commentPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := c.cs.UpdateComment(commentId, commentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
