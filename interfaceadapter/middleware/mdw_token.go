package middleware

import (
	"errors"
	"goclean/interfaceadapter/repository"
	"net/http"
	"time"
)

type MdwToken struct {
	authRepo repository.AuthRepo
	jwtAuth  JwtAuth
}

type JwtAuth interface {
	CreateToken(uid, aud string, nExpiredDay int, signedKey string, now time.Time) (string, error)
	ParseToken(encryptedToken string, repoSignedKeyFunc func(uid, aud string) (string, error)) (string, error)
}

func NewMdwToken(authRepo repository.AuthRepo, jwtAuth JwtAuth) *MdwToken {
	return &MdwToken{
		authRepo: authRepo,
		jwtAuth:  jwtAuth,
	}
}

func (m *MdwToken) HandleFunc(ctrlFunc func(w http.ResponseWriter, r *http.Request, uid string)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// Remove "Bearer "
		if len(authorization) <= 8 {
			ResponseError(w, http.StatusBadRequest, errors.New("Invalid token"))
			return
		}
		token := authorization[7:]

		// parse token, use map claims to avoid float64 to int64 conversion in json decoding
		uid, err := m.jwtAuth.ParseToken(token, m.signedKeyFunc)
		if err != nil {
			ResponseError(w, http.StatusUnauthorized, errors.New("Token expired or invalid"))
			return
		}

		ctrlFunc(w, r, uid)
	})
}

func (m *MdwToken) signedKeyFunc(uid, aud string) (string, error) {
	// Convert jwt.MapClaims to model.Claims
	auth, err := m.authRepo.Get(uid)
	if err != nil {
		return "", err
	}

	signedKey, ok := auth.SignedKeys[aud]
	if !ok {
		return "", nil
	}

	return signedKey.Key, nil
}
