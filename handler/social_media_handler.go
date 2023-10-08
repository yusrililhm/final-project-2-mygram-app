package handler

import (
	"myGram/dto"
	"myGram/pkg/errs"
	"myGram/service/social_media_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediasHandler interface {
	AddSocialMedia(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHndlerImpl struct {
	ss social_media_service.SocialMediaService
}

func NewSocialMediasHandler(socialMediaService social_media_service.SocialMediaService) SocialMediasHandler {
	return &socialMediaHndlerImpl{
		ss: socialMediaService,
	}
}

// AddSocialMedia implements SocialMediasHandler.
func (s *socialMediaHndlerImpl) AddSocialMedia(ctx *gin.Context) {

	socialMediaPayload := &dto.NewSocialMediaRequest{}

	if err := ctx.ShouldBindJSON(socialMediaPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := s.ss.AddSocialMedia(socialMediaPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteSocialMedia implements SocialMediasHandler.
func (s *socialMediaHndlerImpl) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	response, err := s.ss.DeleteSocialMedia(socialMediaId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetSocialMedias implements SocialMediasHandler.
func (s *socialMediaHndlerImpl) GetSocialMedias(ctx *gin.Context) {
	response, err := s.ss.GetSocialMedias()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateSocialMedia implements SocialMediasHandler.
func (s *socialMediaHndlerImpl) UpdateSocialMedia(ctx *gin.Context) {

	socialMediaPayload := &dto.UpdateSocialMediaRequest{}
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if err := ctx.ShouldBindJSON(socialMediaPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := s.ss.UpdateSocialMedia(socialMediaId, socialMediaPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
