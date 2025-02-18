package router

import (
	"RedRock-E-Business/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(h *server.Hertz) {
	// 创建公共路由组
	publicGroup := h.Group("/api/v1")
	{
		// 创建用户相关的子路由组
		userGroup := publicGroup.Group("/user")
		{
			// 注册用户注册路由
			userGroup.POST("/register", handler.UserRegister)
			// 注册用户登录路由
			userGroup.POST("/login", handler.UserLogin)
			// 注册用户信息更新路由
			userGroup.PUT("/update", handler.UpdateUserHandler)
		}
	}
	productGroup := h.Group("/api/v1/products")
	{
		// 注册获取所有商品列表的路由
		productGroup.GET("/", handler.GetAllProductsHandler)
		// 注册根据商品 ID 获取商品详情的路由
		productGroup.GET("/:id", handler.GetProductByIDHandler)
	}
	//commentGroup := publicGroup.Group("/comments")
	//{
	//	commentGroup.PUT("/:comment_id", handler.UpdateCommentHandler)
	//	commentGroup.DELETE("/:comment_id", handler.DeleteCommentHandler)
	//}
}
