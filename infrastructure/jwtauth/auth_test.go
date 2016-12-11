package jwtauth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHashPass(t *testing.T) {
	salt, err := GenSalt()
	assert.Nil(t, err)
	//assert.Equal(t, "123", salt)

	hashedPass1, _ := HashPass("huy123", "So2gdJTHvaY07mT4bQVj6610r00nRYHus6MZ0//PeG0=", "xTARO123x")
	t.Log(hashedPass1, salt)
	assert.Equal(t, "VWrB6FktcxMMQqPTtGJctVfNMw8SThBTRFUhHCUb26I=", hashedPass1)
}

func TestValidatePass(t *testing.T) {
	salt, err := GenSalt()
	assert.Nil(t, err)
	hashedPass, err := HashPass("huy123", salt, "xTARO123x")
	assert.True(t, ValidatePass("huy123", hashedPass, salt, "xTARO123x"))
	assert.False(t, ValidatePass("huy", hashedPass, salt, "xTARO123x"))
}

func TestCreateToken(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Date(2016, 7, 29, 9, 9, 9, 0, loc).Unix()
	claims := jwt.StandardClaims{
		IssuedAt:  now,
		NotBefore: now,
		ExpiresAt: now + int64(time.Duration(7)*24*time.Hour),
	}
	claims.Audience = "default"
	signedKey := "123abc"
	t.Log(GenSalt())
	token, err := CreateToken(claims, signedKey)
	assert.Nil(t, err)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkZWZhdWx0IiwiZXhwIjo2MDQ4MDE0Njk3NTA5NDksImlhdCI6MTQ2OTc1MDk0OSwibmJmIjoxNDY5NzUwOTQ5fQ.8Fbql0HepLC58z0W5ctPraFS-D2IofK6ecuyD_zAtTQ", token)
}
