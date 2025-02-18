package database

import (
	"RedRock-E-Business/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// DB 全局数据库实例
var DB *gorm.DB

// InitMySQL 初始化 MySQL 数据库连接
func InitMySQL() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

// CloseMySQL 关闭 MySQL 数据库连接
func CloseMySQL() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get SQL database instance: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Failed to close database connection: %v", err)
	}
}
func UpdateUser(user *model.User) error {
	return DB.Save(user).Error
}

// GetUserByName 根据用户名获取用户信息
func GetUserByName(name string) (*model.User, error) {
	var user model.User
	result := DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // 用户不存在
		}
		return nil, result.Error // 数据库查询错误
	}
	return &user, nil
}
