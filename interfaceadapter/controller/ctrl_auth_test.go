package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPass(t *testing.T) {
	salt := GenSalt()

	hashedPass1, _ := HashPass("huy123", "So2gdJTHvaY07mT4bQVj6610r00nRYHus6MZ0//PeG0=", "abc")
	t.Log(hashedPass1, salt)
	assert.Equal(t, "atnPKUDYMV/MRIUnjVJVPP/pah4omufxIbFm8H0BrLI=", hashedPass1)
}

func TestValidatePass(t *testing.T) {
	salt := GenSalt()
	hashedPass, _ := HashPass("huy123", salt, "abc")
	assert.True(t, ValidatePass("huy123", hashedPass, salt, "abc"))
	assert.False(t, ValidatePass("huy", hashedPass, salt, "abc"))
}
