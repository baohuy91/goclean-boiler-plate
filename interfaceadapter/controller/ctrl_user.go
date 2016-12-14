package controller

import (
	"net/http"
	"goclean/usecase"
	"github.com/gorilla/mux"
	"errors"
)

type UserCtrl interface {
	GetUser(w http.ResponseWriter, r *http.Request, uid string)
}

func NewUserCtrl(userUsecase usecase.UserUseCase, resp Response) UserCtrl{
	return &userCtrlImpl{
		userUsecase: userUsecase,
		response: resp,
	}
}

type userCtrlImpl struct {
	userUsecase usecase.UserUseCase
	response Response
}

func (c *userCtrlImpl) GetUser(w http.ResponseWriter, r *http.Request, uid string) {
	// Get Uid in query
	vars := mux.Vars(r)
	userId, ok := vars["userId"]

	// Validate request data
	if !ok || userId == "" {
		c.response.Error(w, http.StatusBadRequest, errors.New("missing userId"))
		return
	}

	// Call usecase layer to get user
	userEntity, err := c.userUsecase.GetUser(userId)
	if err!= nil{
		c.response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Convert entity data to the new one that we will response to API
	userPresenter := NewUser(userEntity)

	c.response.Ok(w, userPresenter)
}
