package repository

import (
	"time"
)

type User struct {
	Id string
	Name string
	Email string
	CreatedTime time.Time
	UpdatedTime time.Time
}
