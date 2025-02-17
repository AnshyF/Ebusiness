package dal

import (
	"RedRock-E-Business/database"
	"RedRock-E-Business/model"
)

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckUsernameExists(username string) bool {
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}
