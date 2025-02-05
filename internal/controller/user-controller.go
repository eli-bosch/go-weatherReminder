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

var NewUser models.User

// POST methods
func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)

	u := CreateUser.CreateUser()

	res, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error while marshalling json body")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	//FIX: Add user login method
}

// GET methods
func GetUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, err := json.Marshal(newUsers)
	if err != nil {
		fmt.Println("Error while marshalling users")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing Id")
	}

	userDetails := models.GetUserById(ID)
	res, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println("Error while marshaling json body")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// PUT method
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error parsing Id")
		return
	}

	updateUser := &models.User{}
	utils.ParseBody(r, updateUser)
	userDetails := models.GetUserById(ID)

	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		userDetails.Password = updateUser.Password
	}
	if updateUser.Username != "" {
		userDetails.Username = updateUser.Username
	}

	config.GetDB().Save(&userDetails)

	res, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println("Error while marshalling json body")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Delete method
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
		return
	}

	user := models.DeleteUser(ID)
	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while parsing json body")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
