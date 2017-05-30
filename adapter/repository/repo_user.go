package repository

import (
	"goclean/domain"
	"goclean/usecase"
	"time"
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
		return nil, domain.NewRepoInternalErr(err)
	}

	return toUser(model), err
}

func (r *userRepoImpl) GetByEmail(email string) (*domain.User, error) {
	models := []*UserModel{}
	filter := map[string][]string{"email": {email}}
	err := r.dbGateway.GetPartOfTable(&models, time.Now(), 1, filter)
	if err != nil {
		return nil, domain.NewRepoInternalErr(err)
	}

	if len(models) == 0 {
		return nil, nil
	}

	return toUser(models[0]), nil
}

func (r *userRepoImpl) Create(user domain.User) (string, error) {
	userModel := toUserModel(user)
	id, err := r.dbGateway.Create(userModel)
	if err != nil {
		return nil, domain.NewRepoInternalErr(err)
	}

	return id, nil
}

func toUser(model *UserModel) *domain.User {
	if model == nil {
		return nil
	}

	return &domain.User{
		Id:    model.Id,
		Email: model.Email,
		Name:  model.Name,
	}
}

func toUserModel(user domain.User) *UserModel {
	return &UserModel{
		Id:    "",
		Name:  user.Name,
		Email: user.Email,
	}
}
