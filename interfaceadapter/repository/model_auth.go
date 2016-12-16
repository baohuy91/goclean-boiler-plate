package repository

import "time"

type Auth struct {
	// TODO: add confirmation here
	Uid        string
	SignedKeys map[string]SignedKey
}

type SignedKey struct {
	Name        string
	Key         string
	CreatedTime time.Time
}
