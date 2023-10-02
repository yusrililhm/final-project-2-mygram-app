package photo_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

var (
	AddPhoto func(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error)
)

type photoRepositoryMock struct {
}

func NewPhotoRepositoryMock() PhotoRepository {
	return &photoRepositoryMock{}
}

// AddPhoto implements PhotoRepository.
func (prm *photoRepositoryMock) AddPhoto(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error) {
	return AddPhoto(photoPayload)
}
