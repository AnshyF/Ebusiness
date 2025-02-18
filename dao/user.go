package dao

import (
	"RedRock-E-Business/database"
	"RedRock-E-Business/model"
	"gorm.io/gorm"
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
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil
}
func GetUserByID(id int64) (*model.User, error) {
	var user model.User
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil
}
func UpdateUser(user *model.User) error {
	result := database.DB.Save(user)
	return result.Error
}
