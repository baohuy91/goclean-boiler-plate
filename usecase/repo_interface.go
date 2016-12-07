package usecase

import "goclean/entity"

type UserRepo interface {
	Get(id string) (*entity.User, error)
}
