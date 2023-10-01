package user_repository

import (
	"myGram/entity"
	"myGram/pkg/errs"
)

var (
	Create func(userPayload *entity.User) (int, errs.Error)
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
