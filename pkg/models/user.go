package models

import (
	"gorm.io/gorm"
	"log"
	"modulepath/pkg/config"
)

var db *gorm.DB

type User struct {
	Id      uint64 `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Println("Ошибка при миграции таблицы User:", err)
		return
	}
}

func AddUser(us *User) *User {
	db.Create(&us)
	return us
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(id uint64) (*User, *gorm.DB) {
	var user *User
	db.Find(&user, id)
	return user, db
}

func DeleteUserById(id uint64) *User {
	var user *User
	db.Delete(&user, id)
	return user
}
