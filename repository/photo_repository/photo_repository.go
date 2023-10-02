package photo_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type PhotoRepository interface {
	AddPhoto(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error)
}
