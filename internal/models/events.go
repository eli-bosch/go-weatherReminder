package models

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	LocationID uint `json:"location_id"`

	//Relationship
	User     User     `gorm:"foreignKey"`
	Location Location `gorm:"foreignKey"`
}
