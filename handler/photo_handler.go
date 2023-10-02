package handler

import (
	"myGram/service/photo_service"

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

}

// DeletePhoto implements PhotoHandler.
func (photoHandler *photoHandlerImpl) DeletePhoto(ctx *gin.Context) {

}

// GetPhotos implements PhotoHandler.
func (photoHandler *photoHandlerImpl) GetPhotos(ctx *gin.Context) {

}

// UpdatePhoto implements PhotoHandler.
func (photoHandler *photoHandlerImpl) UpdatePhoto(ctx *gin.Context) {

}
