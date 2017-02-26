package middleware

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomID(t *testing.T) {
	uuid1 := generateRandomID()
	uuid2 := generateRandomID()
	uuid3 := generateRandomID()

	assert.NotEqual(t, uuid1, uuid2)
	assert.NotEqual(t, uuid1, uuid3)
	assert.NotEqual(t, uuid2, uuid3)
}
