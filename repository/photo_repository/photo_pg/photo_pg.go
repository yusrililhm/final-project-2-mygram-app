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

	getUserAndPhotos = `
				SELECT
					p.id,
					p.title,
					p.caption,
					p.photo_url,
					p.user_id,
					p.created_at,
					p.updated_at,
					u.email,
					u.username
				FROM
					photos as p
				LEFT JOIN
					users AS u
				ON
					p.user_id = u.id
				ORDER BY
					p.id
				ASC
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

// GetPhotos implements photo_repository.PhotoRepository.
func (photoRepo *photoRepositoryImpl) GetPhotos() ([]photo_repository.PhotoUserMapped, errs.Error) {

	photosUser := []photo_repository.PhotoUser{}
	rows, err := photoRepo.db.Query(getUserAndPhotos)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	for rows.Next() {
		photoUser := photo_repository.PhotoUser{}

		err = rows.Scan(
			&photoUser.Photo.Id,
			&photoUser.Photo.Title,
			&photoUser.Photo.Caption,
			&photoUser.Photo.PhotoUrl,
			&photoUser.Photo.UserId,
			&photoUser.Photo.CreatedAt,
			&photoUser.Photo.UpdatedAt,
			&photoUser.User.Email,
			&photoUser.User.Username,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("photos not found" + err.Error())
			}
			return nil, errs.NewInternalServerError("something went wrong " + err.Error())
		}

		photosUser = append(photosUser, photoUser)
	}

	result := photo_repository.PhotoUserMapped{}
	return result.HandleMappingPhotoWithUser(photosUser), nil
}
