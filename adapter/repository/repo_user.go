package repository

import (
	"goclean/domain"
	"goclean/usecase"
)

func NewUserRepo() usecase.UserRepo {
	return &userRepoImpl{}
}

type userRepoImpl struct {
	dbGateway DbGateway
}

func (r *userRepoImpl) Get(id string) (*domain.User, error) {
	model := &UserModel{}
	err := r.dbGateway.Get(model, id)
	if err != nil {
		return nil, err
	}

	return toUser(model), nil
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
	userModel := &UserModel{
		Id:    "123",
		Name:  user.Name,
		Email: user.Email,
	}
	return userModel.Id, nil
}

func toUser(model *UserModel) *domain.User {
	return &domain.User{
		Id:    model.Id,
		Email: model.Email,
		Name:  model.Name,
	}
}
