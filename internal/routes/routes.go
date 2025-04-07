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

	//Location Routes
	router.HandleFunc("/location/", controller.CreateLocation).Methods("POST")                 //Creates location from supplied json body
	router.HandleFunc("/location/", controller.GetLocations).Methods("GET")                    //Gets all users - debugging method
	router.HandleFunc("/location/{location_id", controller.GetLocationById).Methods("GET")     //Gets user from user id
	router.HandleFunc("/location/find", controller.FindLocation).Methods("PUT")                //Finds location by json body filled by user
	router.HandleFunc("/location/{location_id}", controller.UpdateLocation).Methods("PUT")     //Updates user from supplied json body
	router.HandleFunc("/loocation/{location_id}", controller.DeleteLocation).Methods("DELETE") //Deletes user and returns their profile

	//Event Routes
	router.HandleFunc("/user/events/", controller.CreateEvent).Methods("POST")             //Creates new Event - needs protections
	router.HandleFunc("/user/events/", controller.GetEvents).Methods("GET")                //Gets all Events for user - needs protections
	router.HandleFunc("/user/events/{event_id}", controller.UpdateEvent).Methods("PUT")    //Update an event - not needed for all intents
	router.HandleFunc("/user/events/{event_id}", controller.DeleteEvent).Methods("DELETE") //Delete an event - needs protections
}
