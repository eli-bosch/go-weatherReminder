package routes

import (
	"github.com/eli-bosch/go-weatherReminder/internal/controller"
	"github.com/gorilla/mux"
)

var RegisterWeatherReminderRoutes = func(router *mux.Router) {
	//User Routes - Add protections and logging
	router.HandleFunc("/user/", controller.CreateUser).Methods("POST")            //Creates user from supplied json body
	router.HandleFunc("/user/", controller.GetUsers).Methods("GET")               //Gets all users - debugging method
	router.HandleFunc("/user/{user_id}", controller.GetUserById).Methods("GET")   //Gets user from user id
	router.HandleFunc("/user/login", controller.LoginUser).Methods("POST")        //Login by sending information in json body
	router.HandleFunc("/user/{user_id}", controller.UpdateUser).Methods("PUT")    //Updates user from supplied json body
	router.HandleFunc("/user/{user_id}", controller.DeleteUser).Methods("DELETE") //Deletes user and returns their profile

	//middleware protection
	//router.HandleFunc("/user/{user_id}", controller.UpdateUser).Methods("PUT").HandlerFunc(middleware.AuthMiddleware) FIX: Add authentication
	//router.HandleFunc("/user/{user_id}", controller.UpdateUser).Methods("DELETE").HandlerFunc(middleware.AuthMiddleware)
}
