package controller

import (
	"net/http"
	"goclean/usecase"
	"github.com/gorilla/mux"
	"errors"
)

type UserCtrl interface {
	GetUser(w http.ResponseWriter, r *http.Request, ctx Context)
}

func NewUserCtrl(userUsecase usecase.UserUseCase) UserCtrl{
	return &userCtrlImpl{
		userUsecase: userUsecase,
	}
}

type userCtrlImpl struct {
	userUsecase usecase.UserUseCase
}

func (c *userCtrlImpl) GetUser(w http.ResponseWriter, r *http.Request, ctx Context) {
	// Get Uid in query
	vars := mux.Vars(r)
	userId, ok := vars["userId"]

	// Validate request data
	if !ok || userId == "" {
		ResponseError(w, http.StatusBadRequest, errors.New("missing userId"))
		return
	}

	// Call usecase layer to get user
	userEntity, err := c.userUsecase.GetUser(userId)
	if err!= nil{
		ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Convert entity data to the new one that we will response to API
	userPresenter := NewUser(userEntity)

	Response(w, userPresenter)
}
