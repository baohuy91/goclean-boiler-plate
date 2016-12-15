package usecase

import (
	"goclean/entity"
	"errors"
)

type UserUseCase interface {
	GetUser(id string) (*entity.User, error)
	RegisterUserByEmail(email string, password string) (*entity.User, error)
}

func NewUserUseCase(userRepo UserRepo, authRepo AuthRepo) UserUseCase {
	return &userUseCaseImpl{
		userRepo:userRepo,
		authRepo:authRepo,
	}
}

type userUseCaseImpl struct {
	userRepo UserRepo
	authRepo AuthRepo
}

// Business logic for getting user will be implemented here
func (u *userUseCaseImpl) GetUser(id string) (*entity.User, error) {
	// Get user from repository & handle error if necessary
	user, err := u.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userUseCaseImpl) RegisterUserByEmail(email string, password string, salt string) (string, error) {
	// Check if email is registered
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Email is already registered")
	}

	// Create User account
	user = &entity.User{
		Email:email,
		Pass: password,
		Salt: salt,
	}
	uid, err := u.userRepo.Create(*user)
	if err != nil {
		return nil, err
	}

	// Create Auth account
	auth := &entity.Auth{
		Uid:uid,
		SignedKeys:map[string]entity.SignedKey{},
	}
	_, err = u.authRepo.Create(*auth)
	if err != nil {
		// TODO: in case create auth fail, remove user record
		return nil, err
	}

	return uid, nil
}
