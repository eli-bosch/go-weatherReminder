package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/eli-bosch/go-weatherReminder/internal/db"
	"github.com/eli-bosch/go-weatherReminder/internal/email"
	"github.com/eli-bosch/go-weatherReminder/internal/models"
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
			runWeatherReminderJob()
		}
	}
}

func runWeatherReminderJob() {
	allEvents := models.GetAllEvents()
	for _, event := range allEvents {
		location := models.GetLocationById(int64(event.LocationID))
		user := models.GetUserById(int64(event.UserID))

		if time.Since(location.UpdatedAt) > 5*time.Minute {
			err := location.UpdateWeatherFields()
			if err != nil {
				fmt.Printf("Failed to fetch weather for ", location)
				continue
			}
		}

		body := location.WeatherDescription + " @ " + strconv.FormatFloat(location.FeelsLike, 'f', 1, 64) + "°F"
		err := email.SendEmailWithSendGrid(user.Email, location.MainWeather, body)
		if err != nil {
			fmt.Printf("Error sending the email to ", user.Email)
		}
	}
}
