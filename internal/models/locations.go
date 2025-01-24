package models

import (
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
