package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"goclean/infrastructure/jwtauth"
	"errors"
	"goclean/usecase"
	"encoding/json"
	"time"
)


// Claims store the data that we stack in the token
type Claims struct {
	SubjectId string `json:"sid,omitempty"`
	Email     string `json:"email,omitempty"`
	jwt.StandardClaims
}

func (c *Claims) InitStandardClaims(day int64) {
	now := time.Now().Unix()
	c.StandardClaims = jwt.StandardClaims{
		IssuedAt:  now,
		NotBefore: now,
		ExpiresAt: now + (day * 24 * 3600 * 1000),
	}
}

func (c *Claims) Init(mapClaims jwt.MapClaims) error {
	assertOk := false
	if sid, ok := mapClaims["sid"]; ok {
		c.SubjectId, assertOk = sid.(string)
		if !assertOk {
			//logrus.Debug("assertOk==false for sid")
		}
	}
	if email, ok := mapClaims["email"]; ok {
		c.Email, assertOk = email.(string)
		if !assertOk {
			//logrus.Debug("assertOk==false for email")
		}
	}
	if aud, ok := mapClaims["aud"]; ok {
		c.Audience, assertOk = aud.(string)
		if !assertOk {
			//logrus.Debug("assertOk==false for aud")
		}
	}
	if nbf, ok := mapClaims["nbf"]; ok {
		notBefore, assertOk := nbf.(json.Number)
		c.NotBefore, _ = notBefore.Int64()
		if !assertOk {
			//logrus.Debug("assertOk==false for nbf, type:", reflect.TypeOf(nbf))
		}
	}
	if iat, ok := mapClaims["iat"]; ok {
		issuedAt, assertOk := iat.(json.Number)
		c.IssuedAt, _ = issuedAt.Int64()
		if !assertOk {
			//logrus.Debug("assertOk==false for iat")
		}
	}
	if exp, ok := mapClaims["exp"]; ok {
		expiresAt, assertOk := exp.(json.Number)
		c.ExpiresAt, _ = expiresAt.Int64()
		if !assertOk {
			//logrus.Debug("assertOk==false for exp")
		}
	}
	if !assertOk {
		return errors.New("Can't convert jwt.MapClaims to model.Claims")
	}
	return nil
}

type MdwToken struct {
	response    Response
	authUseCase usecase.AuthUseCase
}

func NewMdwToken(response Response, authUseCase usecase.AuthUseCase) *MdwToken {
	return &MdwToken{
		response:response,
		authUseCase:authUseCase,
	}
}

func (m *MdwToken) HandleFunc(ctrlFunc func(w http.ResponseWriter, r *http.Request, uid string)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if len(authorization) <= 8 {
			m.response.Error(w, http.StatusBadRequest, errors.New("Invalid token"))
			return
		}

		token := authorization[7:]

		// parse token, use map claims to avoid float64 to int64 conversion in json decoding
		cs, err := jwtauth.ParseToken(token, jwt.MapClaims{}, m.signedKeyFunc)
		if err != nil {
			m.response.Error(w, http.StatusBadRequest, err)
			return
		}
		// Convert jwt.MapClaims to model.Claims
		claims := Claims{}
		err = claims.Init(cs.(jwt.MapClaims))
		if err != nil {
			m.response.Error(w, http.StatusBadRequest, err)
			return
		}

		// authenticate
		if claims.Valid() != nil {
			m.response.Error(w, http.StatusUnauthorized, errors.New("Token expired"))
			return
		}

		ctrlFunc(w, r, claims.SubjectId)
	})
}

func (m *MdwToken)signedKeyFunc(cs jwt.Claims) (interface{}, error) {
	// Convert jwt.MapClaims to model.Claims
	claims := Claims{}
	err := claims.Init(cs.(jwt.MapClaims))
	if err != nil {
		return nil, err
	}

	user, err := m.authUseCase.GetAuth(claims.SubjectId)
	if err != nil {
		return "", err
	}

	signedKey, ok := user.SignedKeys[claims.Audience]
	if !ok {
		return "", nil
	}
	return []byte(signedKey.Key), nil
}
