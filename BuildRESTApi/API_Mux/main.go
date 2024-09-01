package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router from mux library
	router := mux.NewRouter()

	router.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}).Methods("GET");
	router.HandleFunc("/posts", GetPosts).Methods("GET");
	// router.HandleFunc("/posts", AddPost).Methods("POST");


	http.ListenAndServe(":8080", router)
}

