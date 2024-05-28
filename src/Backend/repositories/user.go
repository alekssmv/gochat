package repositories

import (
	"gochat/db"
	"gochat/models"
)

// Возвращает пользователя по username
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	result := db.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// Создает пользователя
func CreateUser(user *models.User) error {
	result := db.Db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}