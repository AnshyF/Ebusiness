// model/comment.go
package model

import (
	"gorm.io/gorm"
)

// Comment 评论模型
type Comment struct {
	gorm.Model
	ProductID int    `json:"product_id"`
	Content   string `json:"content"`
}
