package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/eli-bosch/go-weatherReminder/internal/db"
	"github.com/eli-bosch/go-weatherReminder/internal/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	go startWeatherReminderJob()

	r := mux.NewRouter()
	routes.RegisterWeatherReminderRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

func startWeatherReminderJob() {
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Running reminder job at", time.Now())
			//ADD text api
		}
	}
}
