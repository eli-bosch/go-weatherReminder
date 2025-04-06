package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eli-bosch/go-weatherReminder/config"
	"github.com/eli-bosch/go-weatherReminder/internal/models"
	"github.com/eli-bosch/go-weatherReminder/internal/utils"
	"github.com/gorilla/mux"
)

var NewLocation models.Location

// POST methods
func CreateLocation(w http.ResponseWriter, r *http.Request) {
	CreateLocation := &models.Location{}
	utils.ParseBody(r, CreateLocation)

	l := CreateLocation.CreateLocation()

	res, err := json.Marshal(l)
	if err != nil {
		fmt.Println("Error while marshalling json body")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GET methods
func GetLocations(w http.ResponseWriter, r *http.Request) {
	newLocations := models.GetAllLocations()
	res, err := json.Marshal(newLocations)
	if err != nil {
		fmt.Println("Error while marshalling locations")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLocationById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	locationId := vars["location_id"]
	ID, err := strconv.ParseInt(locationId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
		return
	}

	locationDetails := models.GetLocationById(ID)

	//FIX: Needs to update weather if this gets called (maybe not)

	res, err := json.Marshal(locationDetails)
	if err != nil {
		fmt.Println("Error while marshalling json body")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func FindLocation(w http.ResponseWriter, r *http.Request) {
	//FIX: Finds location by json body
	userLocation := models.Location{}
	utils.ParseBody(r, userLocation)

	if userLocation.City == "" || userLocation.Postal == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}

	location := models.GetLocationByCityAndPostal(userLocation.City, userLocation.Postal)
	if location == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
		return
	}

	res, err := json.Marshal(location)
	if err != nil {
		fmt.Println("Error while parsing Id")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// PUT methods
func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	locationId := vars["location_id"]
	ID, err := strconv.ParseInt(locationId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing Id")
		return
	}

	updateLocation := &models.Location{}
	utils.ParseBody(r, updateLocation)
	locationDetails := models.GetLocationById(ID)

	if updateLocation.Postal != "" {
		locationDetails.Postal = updateLocation.Postal
	}
	if updateLocation.City != "" {
		locationDetails.City = updateLocation.City
	}
	if updateLocation.Region != "" {
		locationDetails.Region = updateLocation.Region
	}
	if updateLocation.Country != "" {
		locationDetails.Country = updateLocation.Country
	}
	if updateLocation.Longitude != 0 {
		locationDetails.Longitude = updateLocation.Longitude
	}
	if updateLocation.Latitude != 0 {
		locationDetails.Latitude = updateLocation.Latitude
	}

	config.GetDB().Save(&locationDetails)

	res, err := json.Marshal(locationDetails)
	if err != nil {
		fmt.Println("Error while marshalling json body")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Delete method
func DeleteLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	locationId := vars["location_id"]
	ID, err := strconv.ParseInt(locationId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
		return
	}

	location := models.DeleteLocation(ID)
	res, err := json.Marshal(location)
	if err != nil {
		fmt.Println("Error while marshalling json body")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
