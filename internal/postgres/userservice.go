package postgres

import (
	"gorm.io/gorm"
	"log"
	"modulepath/internal/models"
)

func init() {
	Connect()
	db = GetDB()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Ошибка при миграции таблицы User:", err)
		return
	}
}

func AddUser(us *models.User) *models.User {
	db.Create(&us)
	return us
}

func GetAllUsers() []models.User {
	var Users []models.User
	result := db.Find(&Users)
	if result.Error != nil {
		log.Println("Ошибка при получении списка пользователей:", result.Error)
	}
	return Users
}

func GetUserById(id uint64) (*models.User, *gorm.DB) {
	var user *models.User
	db.Find(&user, id)
	return user, db
}

func DeleteUserById(id uint64) *models.User {
	var user *models.User
	db.Delete(&user, id)
	return user
}
