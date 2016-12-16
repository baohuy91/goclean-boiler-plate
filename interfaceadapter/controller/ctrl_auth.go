package controller

import (
	"encoding/json"
	"goclean/interfaceadapter/repository"
	"goclean/usecase"
	"io/ioutil"
	"net/http"
)

type AuthCtrl interface {
	LoginByEmail(w http.ResponseWriter, r *http.Request)
	RegisterByMail(w http.ResponseWriter, r *http.Request)
}

func NewAuthCtrl(userUseCase usecase.UserUseCase, authRepo repository.AuthRepo, response Response) AuthCtrl {
	return &authCtrlImpl{
		userUseCase: userUseCase,
		response:    response,
		authRepo:    authRepo,
	}
}

type authCtrlImpl struct {
	userUseCase usecase.UserUseCase
	authRepo    repository.AuthRepo
	response    Response
}

type registerByMailReq struct {
	email string `json:"email"`
	pass  string `json:"pass"`
}

func (c *authCtrlImpl) RegisterByMail(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, err)
		return
	}
	req := registerByMailReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, err)
		return
	}

	// Validate data
	// TODO: validate email format

	// TODO: validate password strength

	// TODO: Generate HashedPass & Salt
	hashPass := ""
	salt := ""

	// Create a user
	userId, err := c.userUseCase.CreateUser()
	if err != nil {
		c.response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Create an auth data
	_, err = c.authRepo.CreateAuthByEmailAndHashPass(userId, req.email, hashPass, salt)
	if err != nil {
		c.response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// TODO: return body response here
	c.response.Ok(w, "")
}

func (c *authCtrlImpl) LoginByEmail(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, err)
		return
	}
	req := registerByMailReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, err)
		return
	}
}
