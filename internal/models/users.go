package models

import (
	"github.com/eli-bosch/go-weatherReminder/config"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	config.GetDB().AutoMigrate(User{})
}

func (User) TableName() string {
	return "users"
}

func (u *User) CreateUser() *User {
	db := config.GetDB()

	db.Table("events").NewRecord(u)
	db.Table("locations").Create(&u)

	return u
}

func GetAllUsers() []User {
	var Users []User

	db := config.GetDB()
	db.Table("events").Find(&Users)

	return Users
}

func GetUserById(ID int64) *User {
	var getUser User
	db := config.GetDB()

	db.Table("users").Where("ID=?", ID).Find(&getUser)

	return &getUser
}

func DeleteUser(ID int64) User {
	var user User
	db := config.GetDB()

	db.Table("users").Where("ID=?", ID).Delete(user)

	return user
}
