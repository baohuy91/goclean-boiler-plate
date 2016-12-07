package repository

import (
	"goclean/entity"
	"goclean/usecase"
)

func NewUserRepo() usecase.UserRepo {
	return &userRepoImpl{}
}

type userRepoImpl struct{}

func (r *userRepoImpl) Get(id string) (*entity.User, error) {
	// TODO: call database
	user := &entity.User{
		Id: id,
	}
	return user, nil
}