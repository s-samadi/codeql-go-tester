// Simple API router handlers

package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "database" of users
var users map[string]UserID = make(map[string]UserID)

func GetUser(rw http.ResponseWriter, r *http.Request) {
	var err error
	defer Recover(&err)

	// get the uuid and validate
	id := mux.Vars(r)["id"]
	log.Println("user: ", id) // <-- logging without sanitization
	err = ValidateID(&id)
	if err != nil {
		HttpErrorResponse(rw, err)
	}

	if resp, ok := users[id]; ok {
		log.Println("user found: ", id) // <-- logging without sanitization
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			HttpErrorResponse(rw, err)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonResp)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}

func CreateUserID(rw http.ResponseWriter, r *http.Request) {
	var err error
	defer Recover(&err)

	// get the payload
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HttpErrorResponse(rw, err)
	}

	// validate the user
	var req UserID
	if err = json.Unmarshal(body, &req); err != nil {
		HttpErrorResponse(rw, err)
	}
	err = req.Validate()
	if err != nil {
		HttpErrorResponse(rw, err)
	}

	// save into our "database"
	id := uuid()
	users[id] = req

	// write json response
	resp := make(map[string]string)
	resp["uuid"] = id
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		HttpErrorResponse(rw, err)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Write(jsonResp)
}

func uuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid = fmt.Sprintf("%04x-%04x-%04x-%04x-%04x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
