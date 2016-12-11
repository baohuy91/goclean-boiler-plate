package repository

import (
	"goclean/usecase"
	"goclean/entity"
)

func NewAuthRepo() usecase.AuthRepo {
	return &authRepoImpl{}
}

type authRepoImpl struct {

}

func (r *authRepoImpl) Get(userId string) (*entity.Auth, error) {
	// TODO: implement here
	return &entity.Auth{}, nil
}
