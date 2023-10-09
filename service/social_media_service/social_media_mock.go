package social_media_service

import (
	"myGram/dto"
	"myGram/pkg/errs"
)

type socialMediaMock struct {
}

var (
	AddSocialMedia    func(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error)
	DeleteSocialMedia func(socialMediaId int) (*dto.GetSocialMediaResponse, errs.Error)
	GetSocialMedias   func() (*dto.GetSocialMediaHttpResponse, errs.Error)
	UpdateSocialMedia func(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error)
)

func NewSocialMediaMock() SocialMediaService {
	return &socialMediaMock{}
}

// AddSocialMedia implements SocialMediaService.
func (s *socialMediaMock) AddSocialMedia(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error) {
	return AddSocialMedia(userId, socialMediaPayload)
}

// DeleteSocialMedia implements SocialMediaService.
func (s *socialMediaMock) DeleteSocialMedia(socialMediaId int) (*dto.GetSocialMediaResponse, errs.Error) {
	return DeleteSocialMedia(socialMediaId)
}

// GetSocialMedias implements SocialMediaService.
func (s *socialMediaMock) GetSocialMedias() (*dto.GetSocialMediaHttpResponse, errs.Error) {
	return GetSocialMedias()
}

// UpdateSocialMedia implements SocialMediaService.
func (s *socialMediaMock) UpdateSocialMedia(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error) {
	return UpdateSocialMedia(socialMediaId, socialMediaPayload)
}
