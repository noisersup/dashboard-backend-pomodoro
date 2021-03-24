package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noisersup/dashboard-backend-pomodoro/handlers"
)

func main() {
	r := mux.NewRouter()

	h := handlers.CreateHandlers()

	r.HandleFunc("/pomodoro", h.GetTimestamp).Methods("GET")

	log.Printf("Starting http server on port :8000...")
	log.Fatal(http.ListenAndServe(":8005", r))
}
