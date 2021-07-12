package services

import (
	"github.com/Tasheem/userServer/dao"
	"github.com/Tasheem/userServer/models"
)

func CreateUser(u models.User) error {
	err := dao.Save(u)
	if err != nil {
		return err
	}

	return nil
}
