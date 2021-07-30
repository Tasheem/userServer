package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}
