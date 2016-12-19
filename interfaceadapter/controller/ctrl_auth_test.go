package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"goclean/interfaceadapter/repository"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

type jwtAuthMock struct {
	createTokenFunc func() (string, error)
	parseTokenFunc  func() (string, error)
}

func (a *jwtAuthMock) CreateToken(uid, aud string, nExpiredDay int, signedKey string, now time.Time) (string, error) {
	if a.createTokenFunc != nil {
		return a.createTokenFunc()
	}
	return "", nil
}

func (a *jwtAuthMock) ParseToken(encryptedToken string, repoSignedKeyFunc func(uid, aud string) (string, error)) (string, error) {
	if a.parseTokenFunc != nil {
		return a.parseTokenFunc()
	}
	return "", nil
}

var tAuthCtrlImpl *authCtrlImpl

func TestMain(m *testing.M) {
	// initialize subject under test
	tAuthCtrlImpl = &authCtrlImpl{
		jwtAuth:     &jwtAuthMock{},
		authRepo:    &authRepoMock{},
		mailManager: &mailManagerMock{},
		userUseCase: &userUseCaseMock{},
	}

	code := m.Run()
	os.Exit(code)
}

func TestAuthCtrlImpl_RegisterByMail_WithNewAuth_ExpectSuccess(t *testing.T) {
	tAuthCtrlImpl = &authCtrlImpl{
		userUseCase: &userUseCaseMock{},
		authRepo:    &authRepoMock{},
	}
	reqData, _ := json.Marshal(registerByMailReq{})
	req, _ := http.NewRequest("POST", "/test", bytes.NewReader(reqData))

	w := httptest.NewRecorder()
	tAuthCtrlImpl.RegisterByMail(w, req)

	assert.Equal(t, 200, w.Code)
	// TODO: unmarshal json & assert response
	//respBody, _ := ioutil.ReadAll(w.Body)
	//assert.Equal(t, "", respBody)
}

func TestAuthCtrlImpl_LoginByEmail_WithExistUser_ExpectValidToken(t *testing.T) {
	tAuthCtrlImpl = &authCtrlImpl{
		authRepo: &authRepoMock{
			getByEmailFunc: func() (*repository.Auth, error) {
				return &repository.Auth{
					SignedKeys: map[string]repository.SignedKey{
						"a123": {},
					},
				}, nil // new user
			},
		},
		jwtAuth: &jwtAuthMock{},
	}
	reqData, _ := json.Marshal(loginByEmailReq{
		aud: "a123",
	})
	req, _ := http.NewRequest("POST", "/test", bytes.NewReader(reqData))

	w := httptest.NewRecorder()
	tAuthCtrlImpl.LoginByEmail(w, req)

	assert.Equal(t, 200, w.Code)
	// TODO: unmarshal json & assert response
}

func TestHashPass(t *testing.T) {
	hashedPass1, _ := HashPass("huy123", "So2gdJTHvaY07mT4bQVj6610r00nRYHus6MZ0//PeG0=", "abc")

	assert.Equal(t, "atnPKUDYMV/MRIUnjVJVPP/pah4omufxIbFm8H0BrLI=", hashedPass1)
}

func TestValidatePass(t *testing.T) {
	salt := GenSalt()

	hashedPass, _ := HashPass("huy123", salt, "abc")

	assert.True(t, ValidatePass("huy123", hashedPass, salt, "abc"))
	assert.False(t, ValidatePass("huy", hashedPass, salt, "abc"))
}
