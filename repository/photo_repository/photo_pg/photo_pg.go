package photo_pg

import (
	"database/sql"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/photo_repository"
	"time"
)

type photoRepositoryImpl struct {
	db *sql.DB
}

const (
	addNewPhotoQuery = `
		INSERT INTO
			photos
				(
					title,
					caption,
					photo_url,
					user_id
				)
		VALUES
				($2, $3, $4, $1)
		RETURNING
				id
	`
)

func NewPhotoRepository(db *sql.DB) photo_repository.PhotoRepository {
	return &photoRepositoryImpl{
		db: db,
	}
}

// AddPhoto implements photo_repository.PhotoRepository.
func (photoRepo *photoRepositoryImpl) AddPhoto(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error) {

	tx, err := photoRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	var id int
	row := tx.QueryRow(addNewPhotoQuery, photoPayload.UserId, photoPayload.Title, photoPayload.Caption, photoPayload.PhotoUrl)

	err = row.Scan(&id)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &dto.PhotoResponse{
		Id:        id,
		Title:     photoPayload.Title,
		Caption:   photoPayload.Caption,
		PhotoUrl:  photoPayload.PhotoUrl,
		UserId:    photoPayload.UserId,
		CreatedAt: time.Now(),
	}, nil
}
