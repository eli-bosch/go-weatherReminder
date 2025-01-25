package models

import (
	"github.com/eli-bosch/go-weatherReminder/config"
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Postal    string  `'gorm:"" json:"postal"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (Location) TableName() string {
	return "locations"
}

func (l *Location) CreateLocation() *Location {
	db := config.GetDB()

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
