package photo_service

import (
	"myGram/dto"
	"myGram/pkg/errs"
)

type photoServiceMock struct {
}

var (
	AddPhoto func(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error)
)

func NewPhotoServiceMock() PhotoService {
	return &photoServiceMock{}
}

// AddPhoto implements PhotoService.
func (psm *photoServiceMock) AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error) {
	return AddPhoto(userId, photoPayload)
}
