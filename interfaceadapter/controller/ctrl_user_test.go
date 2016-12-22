package controller

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"goclean/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserCtrlImpl_GetUser_WithExistUser_ExpectSuccess(t *testing.T) {
	tUserCtrlImpl := &userCtrlImpl{
		userUsecase: &userUseCaseMock{
			getUserFunc: func() (*entity.User, error) {
				return &entity.User{}, nil
			},
		},
	}
	req, _ := http.NewRequest("GET", "/users/u456", nil)

	r := mux.NewRouter()
	r.Path("/users/{userId}").Methods("GET").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tUserCtrlImpl.GetUser(w, r, "u123")
		},
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// TODO: assert data
}

func TestUserCtrlImpl_GetUser_WithNonExistUser_ExpectNotFound(t *testing.T) {
	tUserCtrlImpl := &userCtrlImpl{
		userUsecase: &userUseCaseMock{},
	}
	req, _ := http.NewRequest("GET", "/users/u456", nil)

	r := mux.NewRouter()
	r.Path("/users/{userId}").Methods("GET").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tUserCtrlImpl.GetUser(w, r, "u123")
		},
	)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	// TODO: assert data
}

func TestUserCtrlImpl_GetUser_WithNoUserId_ExpectFail(t *testing.T) {
	tUserCtrlImpl := &userCtrlImpl{
		userUsecase: &userUseCaseMock{},
	}
	req, _ := http.NewRequest("GET", "/test", nil)

	w := httptest.NewRecorder()
	tUserCtrlImpl.GetUser(w, req, "u123")

	assert.Equal(t, 400, w.Code)
	// TODO: assert data
}
