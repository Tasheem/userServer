package main

import (
	"log"
	"net/http"
)

func root(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("User API"))
}

func main() {
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
