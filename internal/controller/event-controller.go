package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eli-bosch/go-weatherReminder/internal/models"
	"github.com/eli-bosch/go-weatherReminder/internal/utils"
)

// POST methods
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	newEvent := &models.Event{}
	utils.ParseBody(r, newEvent)

	e := newEvent.CreateEvent()

	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Error while marshalling json body")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GET methods
func GetEvents(w http.ResponseWriter, r *http.Request) {
	events := models.GetAllEvents()
	res, err := json.Marshal(events)
	if err != nil {
		fmt.Println("Error while marshalling json body")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// PUT methods
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	//Not any reason for this method right now, but implemented for future use
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {

}
