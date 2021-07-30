package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path"
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
	header := req.Header.Get("X-Forwarded-For")
	fmt.Printf("X-Forwarded-For: %v\n", header)

	url := req.URL
	containsQueryString := strings.Contains(url.String(), "?")

	if containsQueryString {
		queryParams := url.Query()

		id := queryParams.Get("id")
		if id != "" {
			fmt.Printf("ID: %s\n", id)
			result, err := services.GetUserByID(id)
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

			res.Header().Add("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			err = json.NewEncoder(res).Encode(result)
			if err != nil {
				fmt.Println("Error Encoding Object to JSON")
				fmt.Println(err)
				http.Error(res, "Error Fetching Users.", http.StatusInternalServerError)
			}
			return
		}

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

		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		err = json.NewEncoder(res).Encode(result)
		if err != nil {
			fmt.Println("Error Encoding Object to JSON")
			fmt.Println(err)
			http.Error(res, "Error Fetching Users.", http.StatusInternalServerError)
		}
	} else { // else get all users
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

func updateUser(res http.ResponseWriter, req *http.Request) {
	userID := path.Base(req.URL.String())
	fmt.Printf("Path Param: %s\n", userID)

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

	fmt.Printf("JSON Object: %v\n", user)

	err := services.UpdateUser(user, userID)
	if err != nil {
		http.Error(res, "Error Creating User.", http.StatusInternalServerError)
		return
	} else {
		res.Write([]byte("User Successfully Updated."))
	}
}

func deleteUser(res http.ResponseWriter, req *http.Request) {

}

func handleUsers(res http.ResponseWriter, req *http.Request) {
	// origin := req.RemoteAddr --> Gets client's hostname correct but gets port incorrect.
	// Client is required to add "Origin" header in request.
	/*origin := req.Header.Get("Origin")
	fmt.Printf("Origin: %v\n", origin)

	// Prevent access to these resources unless client is authServer.
	if origin != "localhost:4000" {
		http.Error(res, "Unauthorized Origin", http.StatusForbidden)
		return
	}*/

	fmt.Printf("Method: %v\n", req.Method)
	if req.Method == "POST" {
		createUser(res, req)
	} else if req.Method == "GET" {
		getUser(res, req)
	} else if req.Method == "PUT" {
		updateUser(res, req)
	} else if req.Method == "DELETE" {
		deleteUser(res, req)
	} else {
		http.Error(res, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}

func root(res http.ResponseWriter, req *http.Request) {
	// origin := req.RemoteAddr --> Gets client's hostname correct but gets port incorrect.
	origin := req.Header.Get("Origin")
	fmt.Printf("Origin: %v\n", origin)

	// Prevent any client from access except for authServer.
	if origin != "localhost:4000" {
		http.Error(res, "Unauthorized Origin", http.StatusForbidden)
		return
	}

	res.Write([]byte("User API"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", root)
	r.HandleFunc("/api/users", handleUsers)
	r.HandleFunc("/api/users/{id}", handleUsers)
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
