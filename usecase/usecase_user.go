package usecase

import (
	"goclean/domain"
)

type UserUseCase interface {
	GetUser(id string) (*domain.User, error)
	CreateUser() (string, error)
}

func NewUserUseCase(userRepo UserRepo) UserUseCase {
	return &userUseCaseImpl{
		userRepo: userRepo,
	}
}

type userUseCaseImpl struct {
	userRepo UserRepo
}

// Business logic for getting user will be implemented here
func (u *userUseCaseImpl) GetUser(id string) (*domain.User, error) {
	// Get user from repository & handle error if necessary
	user, err := u.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userUseCaseImpl) CreateUser() (string, error) {
	// Create User account
	// TODO: Initialize user information here
	user := &domain.User{}
	uid, err := u.userRepo.Create(*user)
	if err != nil {
		return "", err
	}

	return uid, nil
}
