package jwtauth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JwtAuth interface {
	CreateToken(uid, aud string, nExpiredDay int, signedKey string, now time.Time) (string, error)
	ParseToken(encryptedToken string, repoSignedKeyFunc func(uid, aud string) (string, error)) (string, error)
}

type jwtAuthImpl struct{}

func NewJwtAuth() JwtAuth {
	return &jwtAuthImpl{}
}

// Claims store the data that we stack in the token
type Claims struct {
	// Add more claim information if you want here
	jwt.StandardClaims
}

// Create JWT token from uid and aud
// uid to identify which user is this token is given to, or this token can be used to access which user
// aud to identify which kind of client requested token, e.g. iphone 1029d, chrome 88c97,...
// Each aud will have its own signed key, which can be revoked access by user
func (a *jwtAuthImpl) CreateToken(uid, aud string, nExpiredMinute int, signedKey string, now time.Time) (string, error) {
	cs := Claims{
		jwt.StandardClaims{
			Subject:   uid,
			Audience:  aud,
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			ExpiresAt: now.Unix() + int64(nExpiredMinute*60),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cs)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(signedKey))
}

// Parse the token and check if token is valid
// If token is not valid, return error
// Else, return user id
func (a *jwtAuthImpl) ParseToken(encryptedToken string, repoSignedKeyFunc func(uid, aud string) (string, error)) (string, error) {
	parser := jwt.Parser{
		UseJSONNumber: true,
	}

	keyFunc := getSignedKeyFunc(repoSignedKeyFunc)
	// Use default Parse func to use map claims to avoid float64 to int64 conversion in json decoding
	token, err := parser.Parse(encryptedToken, keyFunc)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		// TODO: log suspicious login
		return "", errors.New("Token is invalid")
	}

	// Get necessary information from claims
	uid, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("sub is invalid")
	}

	return uid, nil
}

// Adapter to convert jwt signed key function to controller signed key function
func getSignedKeyFunc(repoSignedKeyFunc func(uid, aud string) (string, error)) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Claim is invalid, not claim map")
		}

		// Get necessary information from claims
		uid, ok := claims["sub"].(string)
		if !ok {
			return nil, errors.New("sub is invalid")
		}

		aud, ok := claims["aud"].(string)
		if !ok {
			return nil, errors.New("aud is invalid")
		}

		signedKey, err := repoSignedKeyFunc(uid, aud)
		if err != nil {
			return "", err
		}
		return []byte(signedKey), nil
	}
}
