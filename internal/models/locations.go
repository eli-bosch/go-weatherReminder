package models

import (
	"github.com/eli-bosch/go-weatherReminder/config"
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

//General Database interactions needed for this project

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

//Private methods

func checkEntryExists(l *Location, db *gorm.DB) *Location {
	var location Location

	return &location
}

//API interactions needed for this project

func addCordinates(l *Location) *Location {

	return l
}
