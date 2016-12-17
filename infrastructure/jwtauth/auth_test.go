package jwtauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJwtAuthImpl_CreateToken(t *testing.T) {
	jwtAuth := jwtAuthImpl{}
	now := time.Date(2016, 12, 16, 23, 30, 30, 0, time.UTC)
	token, err := jwtAuth.CreateToken("uid_123", "ios", 30, "123ab", now)
	assert.Nil(t, err)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpb3MiLCJleHAiOjQwNzM5MzEwMzAsImlhdCI6MTQ4MTkzMTAzMCwibmJmIjoxNDgxOTMxMDMwLCJzdWIiOiJ1aWRfMTIzIn0.v4TYBFoMB9b8bcYJLegVjq_bMwUYHqTbxJrVPlalHLU", token)
}
