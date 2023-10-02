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

	fetchUserByEmail = `
		SELECT
			id, 
			username, 
			email, 
			password, 
			age, 
			created_at, 
			updated_at
		FROM
			users
		WHERE
			email = $1
	`

	fetchUserById = `
		SELECT
			id, 
			username, 
			email, 
			password, 
			age, 
			created_at, 
			updated_at
		FROM
			users
		WHERE
			id = $1
	`

	updateUserQuery = `
		UPDATE 
			users
		SET
			username= $2,
			email= $3,
			updated_at = now()
		WHERE
			id = $1
	`

	deleteUserQuery = `
		DELETE
		FROM
			users
		WHERE
			id = $1
	`
)

func NewUserRepository(db *sql.DB) user_repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// Create implements user_repository.UserRepository.
func (userRepo *userRepositoryImpl) Create(userPayload *entity.User) (int, errs.Error) {
	var id int
	tx, err := userRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return 0, errs.NewInternalServerError("something went wrong")
	}

	err = tx.QueryRow(createUserQuery, userPayload.Username, userPayload.Email, userPayload.Age, userPayload.Password).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, errs.NewInternalServerError("something went wrong")
	}

	return id, nil
}

// Fetch implements user_repository.UserRepository.
func (userRepo *userRepositoryImpl) Fetch(email string) (*entity.User, errs.Error) {

	user := entity.User{}
	err := userRepo.db.QueryRow(fetchUserByEmail, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

// Update implements user_repository.UserRepository.
func (userRepo *userRepositoryImpl) Update(userPayload *entity.User) errs.Error {

	tx, err := userRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong " + err.Error())
	}

	_, err = tx.Exec(updateUserQuery, userPayload.Id, userPayload.Username, userPayload.Email)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong" + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong" + err.Error())
	}

	return nil
}

// FetchById implements user_repository.UserRepository.
func (userRepo *userRepositoryImpl) FetchById(userId int) (*entity.User, errs.Error) {

	user := entity.User{}
	err := userRepo.db.QueryRow(fetchUserById, userId).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

// Delete implements user_repository.UserRepository.
func (userRepo *userRepositoryImpl) Delete(userId int) errs.Error {
	tx, err := userRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deleteUserQuery, userId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
