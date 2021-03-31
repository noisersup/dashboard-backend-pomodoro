package handlers

import (
	"log"
	"net/http"

	"github.com/noisersup/dashboard-backend-pomodoro/database"
	"github.com/noisersup/dashboard-backend-pomodoro/models"
	"github.com/noisersup/dashboard-backend-pomodoro/utils"
)

type PomodoroServer struct {
	db *database.Database
}

func CreateHandlers(db *database.Database) PomodoroServer {
	return PomodoroServer{db}
}

func (srv *PomodoroServer) GetTimestamp(w http.ResponseWriter, r *http.Request) {
	log.Print("GET!") //TODO: remove

	response := models.Response{}

	ts, err := srv.db.GetTimestamp()
	if err != nil { // Database problems [500 code]
		log.Printf("Database error: %s",err) //TODO: Log file

		response.Error = "Database internal error"
		utils.SendResponse(w,response,http.StatusInternalServerError)
		return
	}
	response.Timestamp = ts;
	utils.SendResponse(w, response, http.StatusOK)
}

// func (srv *PomodoroServer) AddTimestamp(w http.ResponseWriter, r *http.Request) {
// 	log.Print("POST!") //TODO: remove

// 	response := models.Response{}

// 	srv.timestamp = int(time.Now().Unix())

// 	response.Timestamp = srv.timestamp
// 	utils.SendResponse(w, response, http.StatusOK)
// }
