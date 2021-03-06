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

func GetUser(username, password string) (models.User, error) {
	return dao.QueryUser(username, password)
}

func GetUserByID(id string) (models.User, error) {
	return dao.QueryUserById(id)
}

func GetUsers() ([]models.User, error) {
	return dao.QueryAll()
}

func UpdateUser(user models.User, userID string) error {
	return dao.Update(userID, user.FirstName, user.LastName)
}

func DeleteUser(userID string) error {
	return dao.Delete(userID)
}