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
// AddComment godoc
// @Summary Add new comment
// @Description Add new comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param dto.NewCommentRequest body dto.NewCommentRequest true "body request for add new comment"
// @Success 201 {object} dto.GetCommentResponse
// @Router /comments [post]
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
// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param commentId path int true "commentId"
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments/{commentId} [delete]
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
// GetComments godoc
// @Summary Get comments
// @Description Get comments
// @Tags Comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments [get]
func (c *commentHandlerImpl) GetComments(ctx *gin.Context) {
	response, err := c.cs.GetComments()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateComment implements CommentHandler.
// UpdateComment godoc
// @Summary Update comment
// @Description Update comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param commentId path int true "commentId"
// @Param dto.UpdateCommentRequest body dto.UpdateCommentRequest true "body request for update comment"
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments/{commentId} [put]
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
