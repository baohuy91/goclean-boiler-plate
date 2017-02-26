package controller

import (
	"goclean/entity"
	"goclean/adapter/repository"
)

// All dependent interface mock will go here

type authRepoMock struct {
	getFunc                          func() (*repository.Auth, error)
	getByEmailFunc                   func() (*repository.Auth, error)
	createAuthByEmailAndHashPassFunc func() (string, error)
	updateFunc                       func() error
	saveSignedKeyFunc                func() error
}

func (r *authRepoMock) Get(userId string) (*repository.Auth, error) {
	if r.getFunc != nil {
		return r.getFunc()
	}
	return nil, nil
}
func (r *authRepoMock) GetByEmail(email string) (*repository.Auth, error) {
	if r.getByEmailFunc != nil {
		return r.getByEmailFunc()
	}
	return nil, nil
}
func (r *authRepoMock) CreateAuthByEmailAndHashPass(uid, email, hashPash, salt string) (string, error) {
	if r.createAuthByEmailAndHashPassFunc != nil {
		return r.createAuthByEmailAndHashPassFunc()
	}
	return "", nil
}
func (r *authRepoMock) Update(auth repository.Auth) error {
	if r.updateFunc != nil {
		return r.updateFunc()
	}
	return nil
}
func (r *authRepoMock) SaveSignedKey(uid, aud, signedKey string) error {
	if r.saveSignedKeyFunc != nil {
		return r.saveSignedKeyFunc()
	}
	return nil
}

type userUseCaseMock struct {
	getUserFunc    func() (*entity.User, error)
	createUserFunc func() (string, error)
}

func (u *userUseCaseMock) GetUser(id string) (*entity.User, error) {
	if u.getUserFunc != nil {
		return u.getUserFunc()
	}
	return nil, nil
}
func (u *userUseCaseMock) CreateUser() (string, error) {
	if u.createUserFunc != nil {
		return u.createUserFunc()
	}
	return "", nil
}
