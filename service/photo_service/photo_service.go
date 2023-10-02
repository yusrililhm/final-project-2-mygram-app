package photo_service

import "myGram/repository/photo_repository"

type PhotoService interface {
	
}

type photoServiceImpl struct {
	pr photo_repository.PhotoRepository
}

func NewPhotoService(photoRepository photo_repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		pr: photoRepository,
	}
}
