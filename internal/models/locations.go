package models

import (
	"errors"
	"fmt"

	"github.com/eli-bosch/go-weatherReminder/config"
	api "github.com/eli-bosch/go-weatherReminder/internal/api/cordinates"
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
}

func (Location) TableName() string {
	return "locations"
}

func (l *Location) CreateLocation() *Location {
	db := config.GetDB()
	exists := checkEntryExists(l, db)

	if exists == nil {
		return exists
	}

	l = addCordinates(l)

	if l == nil {
		return nil
	}

	db.Table("locations").NewRecord(l)
	db.Table("locations").Create(&l)

	return l
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

func checkEntryExists(l *Location, db *gorm.DB) *Location {
	var location Location

	result := db.Table("locations").Where("country=? AND region=? AND city=?", l.Country, l.Region, l.City).First(&location)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //Location is not present in DB
		return l
	}

	if result.Error != nil { //Some other error occured
		return nil
	}

	return &location //Found
}

func addCordinates(l *Location) *Location {
	//FIX: YES
	if l.Longitude == 0 && l.Latitude == 0 {
		cordinateData, err := api.FetchState(l.City, l.Region, l.Country)
		if err != nil {
			fmt.Printf("ERROR: %v", err)
			return nil
		}

		if cordinateData.Country == l.Country && cordinateData.Name == l.City {
			l.Longitude = cordinateData.Longitude
			l.Latitude = cordinateData.Latitude
			return l
		}

		return nil
	}

	return l
}
