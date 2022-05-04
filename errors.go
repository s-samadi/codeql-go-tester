package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpErrorResponse(rw http.ResponseWriter, err error) {
	log.Printf("error: %s", err)
	resp := make(map[string]string)
	resp["message"] = err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("invalid body. Err: %s", err)
	}
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write(jsonResp)
	panic(err)
}

func Recover(errs ...*error) {
	var e *error
	for _, err := range errs {
		e = err
		break
	}
	// handle panic
	if r := recover(); r != nil {
		var errmsg error
		errmsg = r.(error)
		// If error can't bubble up -> Log it
		if e != nil {
			*e = errmsg
		} else {
			log.Printf("%+v", errmsg)
		}
	}
}
