package handlers

import (
	"encoding/json"
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

func (srv *PomodoroServer) SetTimestamp(w http.ResponseWriter, r *http.Request) {
	log.Print("POST!") //TODO: remove

	response := models.Response{}

	timestamp := models.Timestamp{} //TODO: change all to variables

	err := json.NewDecoder(r.Body).Decode(&timestamp)
	if err != nil { // JSON decoding problems [400 code]
		log.Printf("JSON decoding error: %s",err)//TODO: Log file

		response.Error = "Cannot parse json to task object."
		utils.SendResponse(w,response,http.StatusBadRequest)
		return
	}
	if err = srv.db.SetTimestamp(timestamp.Timestamp) ; err != nil{ // Database problems [500 code]
		log.Printf("Database error: %s",err) //TODO: Log file

		response.Error = "Database internal error"
		utils.SendResponse(w,response,http.StatusInternalServerError)
		return
	}
	response.Timestamp = timestamp.Timestamp;
	utils.SendResponse(w, response, http.StatusOK)
}
