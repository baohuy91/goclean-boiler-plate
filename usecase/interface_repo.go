package usecase

import "goclean/entity"

type UserRepo interface {
	Get(id string) (*entity.User, error)
}

type AuthRepo interface {
	Get(userId string) (*entity.Auth, error)
}
