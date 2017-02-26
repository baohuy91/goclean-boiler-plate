package middleware

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

func TestMdwToken_HandleFunc_WithValidToken_ExpectOk(t *testing.T) {
	mdwToken := &MdwToken{
		jwtAuth: &jwtAuthMock{},
	}

	mdwFunc := mdwToken.HandleFunc(func(w http.ResponseWriter, r *http.Request, uid string) {})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	r.Header.Add("Authorization", "Bearer "+"abc")
	mdwFunc.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestMdwToken_HandleFunc_WithJwtParseFail_ExpectFailure(t *testing.T) {
	mdwToken := &MdwToken{
		jwtAuth: &jwtAuthMock{
			parseTokenFunc: func() (string, error) {
				return "", errors.New("")
			},
		},
	}

	mdwFunc := mdwToken.HandleFunc(func(w http.ResponseWriter, r *http.Request, uid string) {})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	r.Header.Add("Authorization", "Bearer "+"abc")
	mdwFunc.ServeHTTP(w, r)

	assert.Equal(t, 401, w.Code)
}

func TestMdwToken_HandleFunc_WithNoToken_ExpectFailure(t *testing.T) {
	mdwToken := &MdwToken{
		jwtAuth: &jwtAuthMock{},
	}

	mdwFunc := mdwToken.HandleFunc(func(w http.ResponseWriter, r *http.Request, uid string) {})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/test", nil)
	mdwFunc.ServeHTTP(w, r)

	assert.Equal(t, 400, w.Code)
}
