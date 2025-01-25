package models

import (
	"fmt"

	"github.com/eli-bosch/go-weatherReminder/config"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) CreateUser() *User {
	db := config.GetDB()
	err := db.Create(&u).Error

	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return nil
	}

	return u
}

func GetAllUsers() []User {
	var Users []User

	db := config.GetDB()
	err := db.Find(&Users).Error
	if err != nil {
		fmt.Printf("Error finding users: %v\n", err)
		return nil
	}

	return Users
}

func GetUserById(ID int64) *User {
	var getUser User
	db := config.GetDB()

	db.Table("users").Where("ID=?", ID).Find(&getUser)

	return &getUser
}

func GetUserByUserName(username string) *User {
	var getUser User
	db := config.GetDB()

	db.Table("users").Where("username=?", username).Find(&getUser)

	return &getUser
}

func GetUserByEmail(email string) *User {
	var getUser User
	db := config.GetDB()

	db.Table("users").Where("email=?", email).Find(&getUser)

	return &getUser
}

func DeleteUser(ID int64) User {
	var user User
	db := config.GetDB()

	db.Table("users").Where("ID=?", ID).Delete(user)

	return user
}
