package repository

import (
	"goclean/entity"
	"goclean/usecase"
	"time"
)

func NewUserRepo() usecase.UserRepo {
	return &userRepoImpl{}
}

type User struct {
	Id string
	Name string
	Email string
	CreatedTime time.Time
	UpdatedTime time.Time
}

type userRepoImpl struct{}

func (r *userRepoImpl) Get(id string) (*entity.User, error) {
	// TODO: call database
	user := &entity.User{
		Id: id,
	}
	return user, nil
}