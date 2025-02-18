package service

import (
	"RedRock-E-Business/database"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"RedRock-E-Business/dao"
	"RedRock-E-Business/middleware"
	"RedRock-E-Business/model"
)

func RegisterUser(req model.RegisterReq) error {
	// 检查用户名是否已存在
	if exists := dao.CheckUsernameExists(req.Username); exists {
		return errors.New("用户名已被注册")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	return dao.CreateUser(&user)
}

func LoginUser(req model.LoginReq) (*model.User, string, error) {
	var user *model.User
	var err error

	// 先尝试按用户名查询
	user, err = dao.GetUserByUsername(req.Identifier)
	if err != nil {
		// 若用户名查询失败，尝试按邮箱查询
		user, err = dao.GetUserByEmail(req.Identifier)
		if err != nil {
			return nil, "", errors.New("用户不存在")
		}
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", errors.New("密码错误")
	}

	// 生成JWT
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, "", errors.New("Token生成失败")
	}

	return user, token, nil
}
func UpdateUserService(user *model.User) error {
	return database.UpdateUser(user)
}
func GetUserByName(Name string) (*model.User, error) {
	return database.GetUserByName(Name)
}
