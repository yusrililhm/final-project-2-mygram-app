package user_repository

import (
	"myGram/entity"
	"myGram/pkg/errs"
)

type UserRepository interface {
	Create(userPayload *entity.User) (int, errs.Error)
	Fetch(email string) (*entity.User, errs.Error)
}
