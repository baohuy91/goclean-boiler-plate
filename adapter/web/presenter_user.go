package web

import "goclean/domain"

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Present user entity to json response
func NewUser(ue *domain.User) *User {
	return &User{
		Id:    ue.Id,
		Name:  ue.Name,
		Email: ue.Email,
	}
}
