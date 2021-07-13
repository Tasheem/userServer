package services

import (
	"fmt"

	"github.com/Tasheem/userServer/dao"
	"github.com/Tasheem/userServer/models"
	"github.com/google/uuid"
)

func CreateUser(u models.User) error {
	u.Id = uuid.New()
	err := dao.Save(u)
	if err != nil {
		return err
	}

	fmt.Printf("User: %v", u)

	return nil
}
