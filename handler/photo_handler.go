package handler

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/service/photo_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	AddPhoto(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHandlerImpl struct {
	ps photo_service.PhotoService
}

func NewPhotoHandler(photoService photo_service.PhotoService) PhotoHandler {
	return &photoHandlerImpl{
		ps: photoService,
	}
}

// AddPhoto implements PhotoHandler.
func (photoHandler *photoHandlerImpl) AddPhoto(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)
	photoPayload := &dto.NewPhotoRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := photoHandler.ps.AddPhoto(user.Id, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeletePhoto implements PhotoHandler.
func (photoHandler *photoHandlerImpl) DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	response, err := photoHandler.ps.DeletePhoto(photoId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetPhotos implements PhotoHandler.
func (photoHandler *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	response, err := photoHandler.ps.GetPhotos()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdatePhoto implements PhotoHandler.
func (photoHandler *photoHandlerImpl) UpdatePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	photoPayload := &dto.NewPhotoRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := photoHandler.ps.UpdatePhoto(photoId, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
