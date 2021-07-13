package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	UserName  string `json:"Username"`
	Password  string `json:"Password"`
}
