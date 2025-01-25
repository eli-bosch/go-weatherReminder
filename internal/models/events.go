package models

import (
	"github.com/eli-bosch/go-weatherReminder/config"
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

func init() {
	config.GetDB().AutoMigrate(Event{})
}

func (Event) TableName() string {
	return "events"
}

func (e *Event) CreateEvent() *Event {
	db := config.GetDB()

	db.Table("events").NewRecord(e)
	db.Table("events").Create(&e)

	return e
}

func GetAllEvents() []Event {
	var events []Event

	db := config.GetDB()
	db.Table("events").Where(&events)

	return events
}

func GetEventById(ID int64) *Event {
	var getEvent Event
	db := config.GetDB()

	db.Table("events").Where("ID=?", ID).Find(&getEvent)

	return &getEvent
}

func GetEventByUserId(ID int64) []Event {
	var getEvents []Event
	db := config.GetDB()

	db.Table("events").Where("user_id=?", ID).Find(&getEvents)

	return getEvents
}

func DeleteEvent(ID int64) Event {
	var event Event
	db := config.GetDB()

	db.Table("events").Where("ID=?", ID).Delete(&event)

	return event
}
