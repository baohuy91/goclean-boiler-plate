package usecase

import "goclean/entity"

type UserRepo interface {
	Get(id string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Create(user entity.User) (string, error)
}
