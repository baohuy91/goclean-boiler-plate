package jwtauth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/sha3"
)

const PW_SALT_SIZE = 32

// TODO: please use lib/email to generate token
func CreateToken(claims jwt.Claims, signedKey string) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(signedKey))
}

func ParseToken(encryptedToken string, claims jwt.Claims, keyFunc func(jwt.Claims) (interface{}, error)) (jwt.Claims, error) {
	parser := jwt.Parser{
		UseJSONNumber: true,
	}
	token, err := parser.ParseWithClaims(encryptedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return keyFunc(token.Claims)
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Token is invalid")
	}
	// TODO: log suspicious login
	return token.Claims, nil
}

// Create a salt with size equal PW_SALT_SIZE
// salt is response as a hex number
func GenSalt() (string, error) {
	b := make([]byte, PW_SALT_SIZE)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func HashPass(pass string, salt string, authSalt string) (string, error) {
	k, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}
	buf := []byte(pass)
	// A MAC with 32 bytes of output has 256-bit security strength -- if you use at least a 32-byte-long key.
	h := make([]byte, 32)
	d := sha3.NewShake256()
	// Write the key into the hash.
	combinedSalt := append(k, []byte(authSalt)...)
	d.Write(combinedSalt)
	// Now write the data.
	d.Write(buf)
	// Read 32 bytes of output from the hash into h.
	d.Read(h)

	return base64.StdEncoding.EncodeToString(h), nil
}

func ValidatePass(pass string, hashedPass string, salt string, authSalt string) bool {
	encoded, err := HashPass(pass, salt, authSalt)
	if err != nil {
		return false
	}
	return encoded == hashedPass
}
