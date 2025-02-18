// handler/product_handler.go
package handler

import (
	"RedRock-E-Business/service"
	"RedRock-E-Business/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// GetAllProductsHandler 处理获取所有商品列表的 HTTP 接口
func GetAllProductsHandler(ctx context.Context, c *app.RequestContext) {
	products, err := service.GetAllProductsService()
	if err != nil {
		utils.SendError(c, 500, "获取商品列表失败: "+err.Error())
		return
	}
	utils.SendSuccess(c, "获取商品列表成功", products)
}

// GetProductByIDHandler 处理根据商品 ID 获取商品详情的 HTTP 接口
func GetProductByIDHandler(ctx context.Context, c *app.RequestContext) {
	// 从请求中获取商品 ID 参数
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// 如果获取 ID 参数失败，返回 400 错误
		utils.SendError(c, 400, "商品 ID 参数错误: "+err.Error())
		return
	}
	// 调用服务层方法根据商品 ID 获取商品详情
	product, err := service.GetProductByIDService(id)
	if err != nil {
		// 如果获取商品详情过程中出现错误，返回 500 错误
		utils.SendError(c, 500, "获取商品详情失败: "+err.Error())
		return
	}
	if product == nil {
		// 如果商品不存在，返回 404 错误
		utils.SendError(c, 404, "商品不存在")
		return
	}
	// 如果成功获取商品详情，返回 200 成功响应
	utils.SendSuccess(c, "获取商品详情成功", product)
}
