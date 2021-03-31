package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func Err(msg string, err error) error{
	return errors.New("["+msg+"] "+err.Error())
}

func SendResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("JSON encoding error: %s", err) //TODO: Log file
	}
}
