package usecase

import "goclean/entity"

type UserRepo interface {
	Get(id string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Create(user entity.User) (string, error)
}

type AuthRepo interface {
	Get(userId string) (*entity.Auth, error)
	Create(auth entity.Auth) (string, error)
}
