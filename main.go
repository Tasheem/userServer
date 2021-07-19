package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func getUser(res http.ResponseWriter, req *http.Request) {
	/* origin := req.RemoteAddr
	fmt.Printf("Origin: %v\n", origin)

	if origin != "localhost:4000" {
		http.Error(res, "Unauthorized Origin", http.StatusForbidden)
		return
	} */

	url := req.URL
	containsQueryString := strings.Contains(url.String(), "?")

	if containsQueryString {
		queryParams := url.Query()
		username := queryParams.Get("Username")
		password := queryParams.Get("Password")

		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Password: %s\n", password)

		result, err := services.GetUser(username, password)
		if err != nil {
			fmt.Println("Error From User Service")
			fmt.Println(err)

			if err.Error() == "user does not exist" {
				http.Error(res, "No Record of User", http.StatusUnauthorized)
			} else {
				http.Error(res, "Error Fetching Users.", http.StatusInternalServerError)
			}
			return
		}

		res.Header().Add("Access-Control-Allow-Origin", "*")
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		err = json.NewEncoder(res).Encode(result)
		if err != nil {
			fmt.Println("Error Encoding Object to JSON")
			fmt.Println(err)
			http.Error(res, "Error Fetching Users.", http.StatusInternalServerError)
		}
	} else {
		users, err := services.GetUsers()
		if err != nil {
			http.Error(res, "Error Fetching Users.", http.StatusInternalServerError)
			return
		}

		res.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(res).Encode(users)
		if err != nil {
			fmt.Println("Error Encoding Object to JSON")
			fmt.Println(err)
			http.Error(res, "Error Fetching Users.", http.StatusInternalServerError)
			return
		}
	}
}

func handleUsers(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Method: %v\n", req.Method);
	if req.Method == "POST" {
		createUser(res, req)
	} else if req.Method == "GET" {
		getUser(res, req)
	}
}

func root(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("User API"))
}

func main() {
	http.HandleFunc("/api", root)
	http.HandleFunc("/api/user", handleUsers)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
