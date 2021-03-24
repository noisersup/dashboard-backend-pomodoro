package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	han "github.com/noisersup/dashboard-backend-pomodoro/handlers"
)

func main() {
	corsPtr := flag.Bool("cors", false, "Enable CORS mode for locally debugging purposes.")
	flag.Parse()

	r := mux.NewRouter()

	h := han.CreateHandlers()

	r.HandleFunc("/pomodoro", h.GetTimestamp).Methods("GET")
	r.HandleFunc("/pomodoro", h.AddTimestamp).Methods("POST")

	var httpHandler http.Handler

	if *corsPtr {
		log.Printf("Warning!!! You are running CORS enabled mode. Do not use it on production")
		headersOk := handlers.AllowedHeaders([]string{"*", "Content-Type"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
		httpHandler = handlers.CORS(originsOk, headersOk, methodsOk)(r)
	} else {
		httpHandler = r
	}

	log.Printf("Starting http server on port :8000...")
	log.Fatal(http.ListenAndServe(":8005", httpHandler))
}
