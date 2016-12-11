package usecase

import "goclean/entity"

type AuthUseCase interface {
	GetAuth(id string) (*entity.Auth, error)
}

func NewAuthUseCase(authRepo AuthRepo) AuthUseCase {
	return &authUseCaseImpl{
		authRepo:authRepo,
	}
}

type authUseCaseImpl struct {
	authRepo AuthRepo
}

// Business logic for getting user will be implemented here
func (u *authUseCaseImpl) GetAuth(id string) (*entity.Auth, error) {
	// Get user from repository & handle error if necessary
	auth, err := u.authRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return auth, err
}