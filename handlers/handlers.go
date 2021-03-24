package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/noisersup/dashboard-backend-pomodoro/models"
	"github.com/noisersup/dashboard-backend-pomodoro/utils"
)

type PomodoroServer struct {
	timestamp int
}

func CreateHandlers() PomodoroServer {
	return PomodoroServer{}
}

func (srv *PomodoroServer) GetTimestamp(w http.ResponseWriter, r *http.Request) {
	log.Print("GET!") //TODO: remove

	response := models.Response{}
	response.Timestamp = srv.timestamp

	utils.SendResponse(w, response, http.StatusOK)
}

func (srv *PomodoroServer) AddTimestamp(w http.ResponseWriter, r *http.Request) {
	log.Print("POST!") //TODO: remove

	response := models.Response{}

	srv.timestamp = int(time.Now().Unix())

	response.Timestamp = srv.timestamp
	utils.SendResponse(w, response, http.StatusOK)
}
