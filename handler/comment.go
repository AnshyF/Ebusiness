// handler/comment_handler.go
package handler

import (
	"RedRock-E-Business/service"
	"RedRock-E-Business/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// CreateCommentHandler 创建评论的 HTTP 接口
func CreateCommentHandler(c *app.RequestContext) {
	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		utils.SendError(c, 400, "商品 ID 参数错误: "+err.Error())
		return
	}
	content := c.PostForm("content")
	comment, err := service.CreateComment(productID, content)
	if err != nil {
		utils.SendError(c, 500, "创建评论失败: "+err.Error())
		return
	}
	utils.SendSuccess(c, "创建评论成功", comment)
}

// GetCommentsByProductIDHandler 根据商品 ID 获取评论列表的 HTTP 接口
func GetCommentsByProductIDHandler(c *app.RequestContext) {
	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		utils.SendError(c, 400, "商品 ID 参数错误: "+err.Error())
		return
	}
	comments, err := service.GetCommentsByProductID(productID)
	if err != nil {
		utils.SendError(c, 500, "获取评论列表失败: "+err.Error())
		return
	}
	utils.SendSuccess(c, "获取评论列表成功", comments)
}

// UpdateCommentHandler 更新评论的 HTTP 接口
func UpdateCommentHandler(c *app.RequestContext) {
	commentIDStr := c.Param("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		utils.SendError(c, 400, "评论 ID 参数错误: "+err.Error())
		return
	}
	content := c.PostForm("content")
	comment, err := service.UpdateComment(commentID, content)
	if err != nil {
		utils.SendError(c, 500, "更新评论失败: "+err.Error())
		return
	}
	utils.SendSuccess(c, "更新评论成功", comment)
}

// DeleteCommentHandler 删除评论的 HTTP 接口
func DeleteCommentHandler(c *app.RequestContext) {
	commentIDStr := c.Param("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		utils.SendError(c, 400, "评论 ID 参数错误: "+err.Error())
		return
	}
	err = service.DeleteComment(commentID)
	if err != nil {
		utils.SendError(c, 500, "删除评论失败: "+err.Error())
		return
	}
	utils.SendSuccess(c, "删除评论成功", nil)
}
