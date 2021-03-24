package handlers

import (
	"log"
	"net/http"
)

type PomodoroServer struct {
	timestamp int
}

func CreateHandlers() PomodoroServer {
	return PomodoroServer{}
}

func (srv *PomodoroServer) GetTimestamp(w http.ResponseWriter, r *http.Request) {
	log.Print("GET!") //TODO: remove

}
