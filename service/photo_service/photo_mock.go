package photo_service

import (
	"myGram/dto"
	"myGram/pkg/errs"
)

type photoServiceMock struct {
}

var (
	AddPhoto    func(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error)
	GetPhotos   func() (*dto.GetPhotoResponse, errs.Error)
	UpdatePhoto func(photoId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error)
)

func NewPhotoServiceMock() PhotoService {
	return &photoServiceMock{}
}

// AddPhoto implements PhotoService.
func (psm *photoServiceMock) AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error) {
	return AddPhoto(userId, photoPayload)
}

// GetPhotos implements PhotoService.
func (psm *photoServiceMock) GetPhotos() (*dto.GetPhotoResponse, errs.Error) {
	return GetPhotos()
}

// UpdatePhoto implements PhotoService.
func (psm *photoServiceMock) UpdatePhoto(photoId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error) {
	return UpdatePhoto(photoId, photoPayload)
}
