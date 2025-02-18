package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:50" json:"username"` // 唯一索引
	Password string `gorm:"not null" json:"-"`                   // 不序列化到JSON
	Email    string `gorm:"uniqueIndex;size:100" json:"email"`   // 唯一邮箱
}

// 注册请求结构体
type RegisterReq struct {
	Username string `json:"username" vd:"len($)>5"` // Hertz参数验证
	Password string `json:"password" vd:"len($)>6"`
	Email    string `json:"email" vd:"email($)"` // 内置邮箱格式验证
}

// 登录请求结构体
type LoginReq struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

// 登录响应结构体
type LoginResp struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

// model/model.go

type UpdateUserReq struct {
	Username string `gorm:"uniqueIndex;size:50" json:"username"` // 唯一索引
	Password string `gorm:"not null" json:"-" vd:"len($)>6"`     // 不序列化到JSON
	Email    string `gorm:"uniqueIndex;size:100" json:"email"`   // 唯一邮箱

}
