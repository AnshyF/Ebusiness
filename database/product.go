// database/product.go
package database

import (
	"RedRock-E-Business/model"
	"gorm.io/gorm"
)

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	result := DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// GetProductByID 根据商品 ID 获取商品详情
func GetProductByID(id int) (*model.Product, error) {
	var product model.Product
	result := DB.First(&product, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &product, nil
}
