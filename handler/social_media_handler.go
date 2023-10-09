package handler

import (
	"myGram/dto"
	"myGram/entity"
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
// AddSocialMedia godoc
// @Summary Add new social media
// @Description Add new social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param dto.NewSocialMediaRequest body dto.NewSocialMediaRequest true "body request for add new social media"
// @Success 201 {object} dto.GetSocialMediaResponse
// @Router /socialmedias [post]
func (s *socialMediaHndlerImpl) AddSocialMedia(ctx *gin.Context) {

	socialMediaPayload := &dto.NewSocialMediaRequest{}
	user := ctx.MustGet("userData").(entity.User)

	if err := ctx.ShouldBindJSON(socialMediaPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := s.ss.AddSocialMedia(user.Id, socialMediaPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteSocialMedia implements SocialMediasHandler.
// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param socialMediaId path int true "socialMediaId"
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socialmedias/{socialMediaId} [delete]
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
// GetSocialMedias godoc
// @Summary Get social medias
// @Description Get social medias
// @Tags Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.GetSocialMediaHttpResponse
// @Router /socialmedias [get]
func (s *socialMediaHndlerImpl) GetSocialMedias(ctx *gin.Context) {
	response, err := s.ss.GetSocialMedias()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateSocialMedia implements SocialMediasHandler.
// UpdateSocialMedia godoc
// @Summary Update social media
// @Description Update social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param socialMediaId path int true "socialMediaId"
// @Param dto.UpdateSocialMediaRequest body dto.UpdateSocialMediaRequest true "body request for update social media"
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socialmedias/{socialMediaId} [put]
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
