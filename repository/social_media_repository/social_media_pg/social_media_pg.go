package social_media_pg

import (
	"database/sql"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/social_media_repository"
)

type socialMediaRepositoryImpl struct {
	db *sql.DB
}

const (
	addSocialMediaQuery = `
		INSERT INTO
			social_media
				(
					name,
					social_media_url
				)
		VALUES
			(
				$1, $2
			)
		RETURNING
			id, name, social_media_url, user_id, created_at
	`
	updateSocialMediaQuery = `
		UPDATE
			comment
		SET
			name = $2,
			social_media_url = $3
		WHERE
			id = $1
	`
	deleteSocialMediaQuery = `
		DELETE FROM
			social_media
		WHERE
			id = $1
	`

	getSocialMediaQuery = `
		SELECT
			s.id,
			s.name,
			s.social_media_url,
			s.user_id,
			s.created_at,
			s.updated_at,
			u.id,
			u.username,
			p.photo_url
		FROM
			social_media AS s
		LEFT JOIN
			users AS u
		ON
			s.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			p.id = s.user_id
		ORDER BY
			s.id
		ASC
	`

	getSocialMediaByIdQuery = `
		SELECT
			s.id,
			s.name,
			s.social_media_url,
			s.user_id,
			s.created_at,
			s.updated_at,
			u.id,
			u.username,
			p.photo_url
		FROM
			social_media AS s
		LEFT JOIN
			users AS u
		ON
			s.user_id = u.id
		LEFT JOIN
			photos AS p
		ON
			p.id = s.user_id
		WHERE
			s.id = $1
	`
)

func NewSocialMediaRepository(db *sql.DB) social_media_repository.SocialMediaRepository {
	return &socialMediaRepositoryImpl{
		db: db,
	}
}

// AddSocialMedia implements social_media_repository.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) AddSocialMedia(socialMediaPayload *entity.SocialMedia) (*dto.NewSocialMediaResponse, errs.Error) {
	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	row := tx.QueryRow(addSocialMediaQuery, socialMediaPayload.Name, socialMediaPayload.SocialMediaUrl)

	var socialMedia dto.NewSocialMediaResponse
	err = row.Scan(
		&socialMedia.Id,
		&socialMedia.Name,
		&socialMedia.SocialMediaUrl,
		&socialMedia.UserId,
		&socialMedia.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &socialMedia, nil
}

// DeleteSocialMedia implements social_media_repository.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) DeleteSocialMedia(socialMediaId int) errs.Error {
	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	_, err = tx.Exec(deleteSocialMediaQuery, socialMediaId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return nil
}

// UpdateSocialMedia implements social_media_repository.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) UpdateSocialMedia(socialMediaId int, socialMediaPayload *entity.SocialMedia) (*dto.SocialMediaUpdateResponse, errs.Error) {
	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	row := tx.QueryRow(updateSocialMediaQuery, socialMediaPayload.Name, socialMediaPayload.SocialMediaUrl)

	var socialMedia dto.SocialMediaUpdateResponse
	err = row.Scan(
		&socialMedia.Id,
		&socialMedia.Name,
		&socialMedia.SocialMediaUrl,
		&socialMedia.UserId,
		&socialMedia.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &socialMedia, nil
}

// GetSocialMediaById implements social_media_repository.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) GetSocialMediaById(socialMediaId int) (*dto.GetSocialMedia, errs.Error) {
	row := s.db.QueryRow(getSocialMediaByIdQuery, socialMediaId)

	var socialMedia social_media_repository.SocialMediaUserPhoto
	err := row.Scan(
		&socialMedia.SocialMedia.Id,
		&socialMedia.SocialMedia.Name,
		&socialMedia.SocialMedia.SocialMediaUrl,
		&socialMedia.SocialMedia.UserId,
		&socialMedia.SocialMedia.CreatedAt,
		&socialMedia.SocialMedia.UpdatedAt,
		&socialMedia.User.Id,
		&socialMedia.User.Username,
		&socialMedia.Photo.PhotoUrl,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("social media not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	result := social_media_repository.SocialMediaUserPhotoMapped{}
	return result.HandleMappingSocialMediaWithUserAndPhotoById(socialMedia), nil
}

// GetSocialMedias implements social_media_repository.SocialMediaRepository.
func (s *socialMediaRepositoryImpl) GetSocialMedias() ([]*dto.GetSocialMedia, errs.Error) {

	rows, err := s.db.Query(getSocialMediaQuery)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var socialMedias []social_media_repository.SocialMediaUserPhoto

	for rows.Next() {
		var socialMedia social_media_repository.SocialMediaUserPhoto
		err = rows.Scan(
			&socialMedia.SocialMedia.Id,
			&socialMedia.SocialMedia.Name,
			&socialMedia.SocialMedia.SocialMediaUrl,
			&socialMedia.SocialMedia.UserId,
			&socialMedia.SocialMedia.CreatedAt,
			&socialMedia.SocialMedia.UpdatedAt,
			&socialMedia.User.Id,
			&socialMedia.User.Username,
			&socialMedia.Photo.PhotoUrl,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		socialMedias = append(socialMedias, socialMedia)
	}

	result := social_media_repository.SocialMediaUserPhotoMapped{}
	return result.HandleMappingSocialMediaWithUserAndPhoto(socialMedias), nil
}
