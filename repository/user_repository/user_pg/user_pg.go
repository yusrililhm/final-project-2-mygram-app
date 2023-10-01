package user_pg

import (
	"database/sql"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/user_repository"
)

type userRepositoryImpl struct {
	db *sql.DB
}

const (
	createUserQuery = `
		INSERT INTO 
			users (username, email, age, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`
)

func NewUserRepository(db *sql.DB) user_repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// Create implements user_repository.UserRepository.
func (userRepo *userRepositoryImpl) Create(userPayload *entity.User) (*int, errs.Error) {
	var id int
	tx, err := userRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = tx.QueryRow(createUserQuery, userPayload.Username, userPayload.Email, userPayload.Age, userPayload.Password).Scan(&id)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wron")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wron")
	}

	return &id, nil
}
