package usecase

import "goclean/domain"

type UserRepo interface {
	Get(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Create(user domain.User) (string, error)
}
