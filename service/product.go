// service/product_service.go
package service

import (
	"RedRock-E-Business/database"
	"RedRock-E-Business/model"
)

// GetAllProductsService 获取所有商品列表的业务逻辑
func GetAllProductsService() ([]model.Product, error) {
	return database.GetAllProducts()
}

// GetProductByIDService 根据商品 ID 获取商品详情的业务逻辑
func GetProductByIDService(id int) (*model.Product, error) {
	return database.GetProductByID(id)
}
