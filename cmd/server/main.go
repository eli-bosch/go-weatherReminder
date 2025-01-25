package main

import (
	"log"
	"net/http"

	_ "github.com/eli-bosch/go-weatherReminder/internal/db"
	"github.com/eli-bosch/go-weatherReminder/internal/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterWeatherReminderRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
