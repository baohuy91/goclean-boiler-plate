package entity

import "time"

type Auth struct{
	Uid string
	SignedKeys map[string] SignedKey
}

type SignedKey struct {
	Name        string    `gorethink:"name"`
	Key         string    `gorethink:"key"`
	CreatedTime time.Time `gorethink:"createdTime"`
}
