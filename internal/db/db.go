package db

import (
	"fmt"

	"github.com/eli-bosch/go-weatherReminder/config"
	"github.com/eli-bosch/go-weatherReminder/internal/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {

	db := config.Connect()

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Location{})
	db.AutoMigrate(models.Event{})

	fmt.Println("Database is connected...")
}
