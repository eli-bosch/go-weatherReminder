package db

import (
	"fmt"

	"github.com/eli-bosch/go-weatherReminder/internal/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	config.Connect()
	fmt.Println("Database is connected...")
}
