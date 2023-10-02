package photo_pg

import (
	"database/sql"
	"myGram/repository/photo_repository"
)

type photoRepositoryImpl struct {
	db *sql.DB
}

func NewPhotoRepository(db *sql.DB) photo_repository.PhotoRepository {
	return &photoRepositoryImpl{
		db: db,
	}
}
