package models

import (
	"fmt"

	"github.com/eli-bosch/go-weatherReminder/config"
	cordinates "github.com/eli-bosch/go-weatherReminder/internal/api/cordinates"
	weather "github.com/eli-bosch/go-weatherReminder/internal/api/weather"
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Postal             string  `'gorm:"" json:"postal"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	Country            string  `json:"country"`
	Longitude          float64 `json:"longitude"`
	Latitude           float64 `json:"latitude"`
	MainWeather        string  `gorm:"main_weather"`
	WeatherDescription string  `gorm:"description_weather"`
	FeelsLike          float64 `gorm:feels_like`
}

func (Location) TableName() string {
	return "locations"
}

//General Database interactions needed for this project

func (l *Location) CreateLocation() *Location {
	db := config.GetDB()
	exists := GetLocationByCityAndPostal(l.City, l.Postal)

	if exists != nil {
		return exists
	}

	l = addCordinates(l)
	if l == nil {
		return nil
	}

	l = addWeather(l)
	if l == nil {
		return nil
	}

	db.Table("locations").NewRecord(l)
	db.Table("locations").Create(&l)

	return l
}

func (l *Location) UpdateWeatherFields() error {
	db := config.GetDB()

	var existing Location
	if err := db.First(&existing, l.ID).Error; err != nil {
		fmt.Println("ERROR: Location not found for weather update")
		return err
	}

	l = addWeather(l)
	if l == nil {
		return fmt.Errorf("ERROR: failed to fetch weather")
	}

	// Apply the updated weather data to the existing record
	db.Model(&existing).Updates(map[string]interface{}{
		"latitude":            l.Latitude,
		"longitude":           l.Longitude,
		"main_weather":        l.MainWeather,
		"description_weather": l.WeatherDescription,
		"feels_like":          l.FeelsLike,
	})

	return nil
}

func GetAllLocations() []Location {
	var Locations []Location

	db := config.GetDB()
	db.Table("locations").Find(&Locations)

	return Locations
}

func GetLocationByCityAndPostal(city string, postal string) *Location {
	var getLocation Location

	db := config.GetDB()
	db.Table("locations").Where("city=? AND postal=?", city, postal).Find(&getLocation)

	return &getLocation
}

func GetLocationById(ID int64) *Location {
	var getLocation Location
	db := config.GetDB()

	db.Table("locations").Where("ID=?", ID).Find(&getLocation)

	return &getLocation
}

func DeleteLocation(ID int64) Location {
	var location Location
	db := config.GetDB()

	db.Table("locations").Where("ID=?", ID).Delete(location)

	return location
}

//API interactions needed for this project

func addCordinates(l *Location) *Location {
	cord, err := cordinates.FetchState(l.City, l.Region, l.Country)
	if err != nil {
		fmt.Println("ERROR FETCHING CORDINATES: ", err)
		return nil
	}

	l.Latitude = cord.Latitude
	l.Longitude = cord.Longitude

	return l
}

func addWeather(l *Location) *Location {
	curWeather, err := weather.FetchWeather(l.Longitude, l.Latitude)
	if err != nil {
		fmt.Println("ERROR FETCHING CORDINATES: ", err)
		return nil
	}

	l.MainWeather = curWeather.Weather[0].Main
	l.WeatherDescription = curWeather.Weather[0].Description
	l.FeelsLike = curWeather.Main.FeelsLike

	return l
}
