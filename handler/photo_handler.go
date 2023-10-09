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
// AddPhoto godoc
// @Summary Add new photo
// @Description Add new photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param dto.NewPhotoRequest body dto.NewPhotoRequest true "body request for add new photo"
// @Success 201 {object} dto.GetPhotoResponse
// @Router /photos [post]
func (p *photoHandlerImpl) AddPhoto(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)
	photoPayload := &dto.NewPhotoRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := p.ps.AddPhoto(user.Id, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeletePhoto implements PhotoHandler.
// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param photoId path int true "photoId"
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos/{photoId} [delete]
func (p *photoHandlerImpl) DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	response, err := p.ps.DeletePhoto(photoId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetPhotos implements PhotoHandler.
// GetPhotos godoc
// @Summary Get photos
// @Description Get photos
// @Tags Photos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos [get]
func (p *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	response, err := p.ps.GetPhotos()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdatePhoto implements PhotoHandler.
// UpdatePhoto godoc
// @Summary Update photo
// @Description Update photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param photoId path int true "photoId"
// @Param dto.PhotoUpdateRequest body dto.PhotoUpdateRequest true "body request for update photo"
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos/{photoId} [put]
func (p *photoHandlerImpl) UpdatePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	photoPayload := &dto.PhotoUpdateRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := p.ps.UpdatePhoto(photoId, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
