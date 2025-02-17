package handler

import (
	"context"

	"RedRock-E-Business/model"
	"RedRock-E-Business/service"
	"RedRock-E-Business/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req model.RegisterReq
	if err := c.BindAndValidate(&req); err != nil {
		utils.SendError(c, 400, "参数错误: "+err.Error())
		return
	}

	if err := service.RegisterUser(req); err != nil {
		utils.SendError(c, 500, "注册失败: "+err.Error())
		return
	}

	utils.SendSuccess(c, "注册成功", nil)
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var req model.LoginReq
	if err := c.BindAndValidate(&req); err != nil {
		utils.SendError(c, 400, "参数错误")
		return
	}

	user, token, err := service.LoginUser(req)
	if err != nil {
		utils.SendError(c, 401, "登录失败: "+err.Error())
		return
	}

	utils.SendSuccess(c, "登录成功", model.LoginResp{
		UserID: user.ID,
		Token:  token,
	})
}
func UpdateUserHandler(ctx context.Context, c *app.RequestContext) {
	var req model.UpdateUserReq
	if err := c.BindAndValidate(&req); err != nil {
		utils.SendError(c, 400, "参数错误: "+err.Error())
		return
	}

	// 根据请求中的用户 ID 获取用户信息
	user, err := service.GetUserByName(req.Username)
	if err != nil {
		utils.SendError(c, 404, "用户不存在: "+err.Error())
		return
	}

	// 更新用户信息
	user.Username = req.Username
	user.Email = req.Email
	// 可以根据需要添加更多字段的更新

	if err := service.UpdateUserService(user); err != nil {
		utils.SendError(c, 500, "更新用户信息失败: "+err.Error())
		return
	}

	utils.SendSuccess(c, "用户信息更新成功", nil)
}
