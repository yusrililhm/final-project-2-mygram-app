package user_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

var (
	Create    func(userPayload *entity.User) (*dto.UserResponse, errs.Error)
	Fetch     func(email string) (*entity.User, errs.Error)
	Update    func(userPayload *entity.User) (*dto.UserUpdateResponse, errs.Error)
	FetchById func(userId int) (*entity.User, errs.Error)
	Delete    func(userId int) errs.Error
)

type userRepositoryMock struct {
}

func NewUserRepositoryMock() UserRepository {
	return &userRepositoryMock{}
}

// Create implements UserRepository.
func (urm *userRepositoryMock) Create(userPayload *entity.User) (*dto.UserResponse, errs.Error) {
	return Create(userPayload)
}

// Fetch implements UserRepository.
func (urm *userRepositoryMock) Fetch(email string) (*entity.User, errs.Error) {
	return Fetch(email)
}

// Update implements UserRepository.
func (urm *userRepositoryMock) Update(userPayload *entity.User) (*dto.UserUpdateResponse, errs.Error) {
	return Update((userPayload))
}

// FetchById implements UserRepository.
func (urm *userRepositoryMock) FetchById(userId int) (*entity.User, errs.Error) {
	return FetchById(userId)
}

// Delete implements UserRepository.
func (urm *userRepositoryMock) Delete(userId int) errs.Error {
	return Delete(userId)
}
