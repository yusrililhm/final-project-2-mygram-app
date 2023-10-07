package photo_service

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/photo_repository"
	"net/http"
)

type PhotoService interface {
	AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error)
	GetPhotos() (*dto.GetPhotoResponse, errs.Error)
	UpdatePhoto(photoId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error)
	DeletePhoto(photoId int) (*dto.GetPhotoResponse, errs.Error)
}

type photoServiceImpl struct {
	pr photo_repository.PhotoRepository
}

func NewPhotoService(photoRepository photo_repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		pr: photoRepository,
	}
}

// AddPhoto implements PhotoService.
func (photoService *photoServiceImpl) AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error) {

	err := helper.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	photo := &entity.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
		UserId:   userId,
	}

	response, err := photoService.pr.AddPhoto(photo)

	if err != nil {
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusCreated,
		Message:    "Add new photo successfully",
		Data:       response,
	}, nil
}

// GetPhotos implements PhotoService.
func (photoService *photoServiceImpl) GetPhotos() (*dto.GetPhotoResponse, errs.Error) {

	result, err := photoService.pr.GetPhotos()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "fetch photos successfully",
		Data:       result,
	}, nil
}

// UpdatePhoto implements PhotoService.
func (photoService *photoServiceImpl) UpdatePhoto(photoId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, errs.Error) {

	err := helper.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	photo := &entity.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
	}

	err = photoService.pr.UpdatePhoto(photoId, photo)

	if err != nil {
		return nil, err
	}

	response, err := photoService.pr.GetPhotoId(photoId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "photo successfully updated",
		Data: dto.PhotoUpdateResponse{
			Id:        response.Id,
			Title:     response.Title,
			Caption:   response.Caption,
			PhotoUrl:  response.PhotoUrl,
			UserId:    response.UserId,
			UpdatedAt: response.UpdatedAt,
		},
	}, nil
}

// DeletePhoto implements PhotoService.
func (photoService *photoServiceImpl) DeletePhoto(photoId int) (*dto.GetPhotoResponse, errs.Error) {

	err := photoService.pr.DeletePhoto(photoId)

	if err != nil {
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "Your photo has been successfully deleted",
		Data:       nil,
	}, nil
}
