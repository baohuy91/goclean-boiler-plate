package jwtauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJwtAuthImpl_CreateToken(t *testing.T) {
	jwtAuth := jwtAuthImpl{}
	now := time.Date(2016, 12, 16, 23, 30, 30, 0, time.UTC)

	token, err := jwtAuth.CreateToken("uid_123", "ios", 20*365*24*60, "123ab", now)

	assert.Nil(t, err)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpb3MiLCJleHAiOjIxMTI2NTEwMzAsImlhdCI6MTQ4MTkzMTAzMCwibmJmIjoxNDgxOTMxMDMwLCJzdWIiOiJ1aWRfMTIzIn0.lKNyxYDwCTyVsph4zvRnO8o6cByItAd5ESsUY8KZUVA", token)
}

func TestJwtAuthImpl_ParseToken(t *testing.T) {
	jwtAuth := jwtAuthImpl{}
	encryptedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpb3MiLCJleHAiOjIxMTI2NTEwMzAsImlhdCI6MTQ4MTkzMTAzMCwibmJmIjoxNDgxOTMxMDMwLCJzdWIiOiJ1aWRfMTIzIn0.lKNyxYDwCTyVsph4zvRnO8o6cByItAd5ESsUY8KZUVA"

	uid, err := jwtAuth.ParseToken(encryptedToken, func(uid, aud string) (string, error) {
		return "123ab", nil
	})

	assert.NoError(t, err)
	assert.Equal(t, "uid_123", uid)
}
