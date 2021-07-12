package dao

import (
	"database/sql"
	"fmt"

	"github.com/Tasheem/userServer/models"
)

var (
	username = "root"
	password = "colts1810"
	address  = "127.0.0.1:3306"
)

func createDBIfDoesNotExist() error {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/", username, password, address)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS BookStore;")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec("USE BookStore")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(" +
		"id varchar(36) NOT NULL," +
		"first_name varchar(100)," +
		"last_name varchar(100)," +
		"username varchar(100)," +
		"password varchar(30)," +
		"PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Save(u models.User) error {
	err := createDBIfDoesNotExist()
	if err != nil {
		fmt.Println(err)
		return err
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/BookStore", username, password, address)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("dao->Save: Error Opening SQL Connection statement")
		fmt.Println(err)
		return err
	}
	defer db.Close()

	insert := fmt.Sprintf("INSERT INTO users VALUES (\"%s\", \"%s\", \"%s\", \"%s\", \"%s\");",
		u.Id.String(), u.FirstName, u.LastName, u.UserName, u.Password)

	_, err = db.Exec(insert)
	if err != nil {
		fmt.Println("dao->Save: Error With INSERT statement")
		fmt.Println(err)
		fmt.Printf("Insert Statement: %s", insert)
		return err
	}

	return nil
}
