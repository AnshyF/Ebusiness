// service/comment.go
package service

import (
	"RedRock-E-Business/database"
)

// Comment 评论结构体
type Comment struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Content   string `json:"content"`
}

// CreateComment 创建评论
func CreateComment(productID int, content string) (*Comment, error) {
	comment := &Comment{
		ProductID: productID,
		Content:   content,
	}
	result := database.DB.Create(comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

// GetCommentsByProductID 根据商品 ID 获取评论列表
func GetCommentsByProductID(productID int) ([]*Comment, error) {
	var comments []*Comment
	result := database.DB.Where("product_id = ?", productID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

// UpdateComment 更新评论
func UpdateComment(commentID int, content string) (*Comment, error) {
	var comment Comment
	result := database.DB.First(&comment, commentID)
	if result.Error != nil {
		return nil, result.Error
	}
	comment.Content = content
	result = database.DB.Save(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

// DeleteComment 删除评论
func DeleteComment(commentID int) error {
	result := database.DB.Delete(&Comment{}, commentID)
	return result.Error
}
