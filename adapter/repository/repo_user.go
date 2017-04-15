package repository

import (
	"goclean/domain"
	"goclean/usecase"
)

func NewUserRepo() usecase.UserRepo {
	return &userRepoImpl{}
}

type User struct {
	Id    string
	Name  string
	Email string
	Pass  string
	Salt  string
	CommonModelImpl
}

type userRepoImpl struct{}

func (r *userRepoImpl) Get(id string) (*domain.User, error) {
	// TODO: call database
	user := &domain.User{
		Id: id,
	}
	return user, nil
}

func (r *userRepoImpl) GetByEmail(email string) (*domain.User, error) {
	// TODO: call database
	user := &domain.User{
		Id:    "123",
		Email: email,
	}
	return user, nil
}

func (r *userRepoImpl) Create(user domain.User) (string, error) {
	// TODO: call database
	userModel := &User{
		Id:    "123",
		Name:  user.Name,
		Email: user.Email,
		Pass:  user.HashPass,
		Salt:  user.Salt,
	}
	return userModel.Id, nil
}
