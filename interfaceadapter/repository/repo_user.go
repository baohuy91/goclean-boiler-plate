package repository

import (
	"goclean/entity"
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
	BaseModelImpl
}

type userRepoImpl struct{}

func (r *userRepoImpl) Get(id string) (*entity.User, error) {
	// TODO: call database
	user := &entity.User{
		Id: id,
	}
	return user, nil
}

func (r *userRepoImpl) GetByEmail(email string) (*entity.User, error) {
	// TODO: call database
	user := &entity.User{
		Id: "123",
		Email: email,
	}
	return user, nil
}

func (r *userRepoImpl) Create(user entity.User) (string, error) {
	// TODO: call database
	userModel := &User{
		Id: "123",
		Name: user.Name,
		Email: user.Email,
		Pass: user.Pass,
		Salt: user.Salt,
	}
	return userModel.Id, nil
}
