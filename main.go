package main

// Main API server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("running: localhost:8000")

	router := mux.NewRouter()
	router.Handle("/user/{id}", AuthorizationMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
	router.Handle("/user", http.HandlerFunc(CreateUserID)).Methods("POST")
	http.Handle("/", router)

	http.ListenAndServe(":8000", router)
}
