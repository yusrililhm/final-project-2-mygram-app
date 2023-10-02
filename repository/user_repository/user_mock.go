package user_repository

import (
	"myGram/entity"
	"myGram/pkg/errs"
)

var (
	Create    func(userPayload *entity.User) (int, errs.Error)
	Fetch     func(email string) (*entity.User, errs.Error)
	Update    func(userPayload *entity.User) errs.Error
	FetchById func(userId int) (*entity.User, errs.Error)
)

type userRepositoryMock struct {
}

func NewUserRepositoryMock() UserRepository {
	return &userRepositoryMock{}
}

// Create implements UserRepository.
func (urm *userRepositoryMock) Create(userPayload *entity.User) (int, errs.Error) {
	return Create(userPayload)
}

// Fetch implements UserRepository.
func (urm *userRepositoryMock) Fetch(email string) (*entity.User, errs.Error) {
	return Fetch(email)
}

// Update implements UserRepository.
func (urm *userRepositoryMock) Update(userPayload *entity.User) errs.Error {
	return Update((userPayload))
}

// FetchById implements UserRepository.
func (urm *userRepositoryMock) FetchById(userId int) (*entity.User, errs.Error) {
	return FetchById(userId)
}
