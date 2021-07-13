package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tasheem/userServer/models"
	"github.com/Tasheem/userServer/services"
	_ "github.com/go-sql-driver/mysql"
)

func createUser(res http.ResponseWriter, req *http.Request) {
	var user models.User

	if req.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(req.Body).Decode(&user)

		if err != nil {
			http.Error(res, "Invalid JSON", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(res, "Invalid Media Type", http.StatusUnsupportedMediaType)
		return
	}

	go func() {
		fmt.Printf("JSON Object: %v\n", user)
	}()

	err := services.CreateUser(user)
	if err != nil {
		http.Error(res, "Error Creating User.", http.StatusInternalServerError)
		return
	} else {
		res.Write([]byte("User Successfully Created."))
	}
}

func user(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		createUser(res, req)
	}
}

func root(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("User API"))
}

func main() {
	http.HandleFunc("/api", root)
	http.HandleFunc("/api/user", user)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
