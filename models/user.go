package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}
