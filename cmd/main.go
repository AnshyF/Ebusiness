package main

import (
	router "RedRock-E-Business/api"
	config "RedRock-E-Business/configs"
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func main() {
	// -------------------- 1. 初始化配置 --------------------
	if err := config.Init("./configs"); err != nil {
		panic(fmt.Sprintf("配置加载失败: %v", err))
	}
	hlog.Info("配置文件加载成功")

	// -------------------- 2. 初始化基础依赖 --------------------
	// 初始化 MySQL
	//if err := database.InitMySQL(); err != nil {
	//	panic(fmt.Sprintf("MySQL连接失败: %v", err))
	//}
	//defer database.CloseMySQL() // 程序退出时关闭连接
	//hlog.Info("MySQL连接成功")

	// 初始化 Redis
	//if err := redis.InitRedis(); err != nil {
	//	panic(fmt.Sprintf("Redis连接失败: %v", err))
	//}
	//defer redis.CloseRedis()
	//hlog.Info("Redis连接成功")

	// -------------------- 3. 创建 Hertz 实例 --------------------
	h := server.New(
		server.WithHostPorts(config.Conf.Server.Port), // 监听端口 // 使用标准网络库
		server.WithHandleMethodNotAllowed(true),       // 允许处理 405 错误
		server.WithDisablePrintRoute(false),           // 打印注册的路由信息（调试用）
	)

	// -------------------- 4. 注册全局中间件 --------------------
	h.Use(
		// 添加 CORS 跨域支持
		func(ctx context.Context, c *app.RequestContext) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Next(ctx)
		},
		// 添加请求日志中间件
		func(ctx context.Context, c *app.RequestContext) {
			hlog.CtxInfof(ctx, "Request: %s %s", c.Method(), c.Path())
			c.Next(ctx)
		},
	)

	// -------------------- 5. 注册路由 --------------------
	router.RegisterRoutes(h) // 将所有路由绑定到 Hertz 实例

	// -------------------- 6. 启动服务 --------------------
	hlog.Info("服务启动，监听端口: ", config.Conf.Server.Port)
	h.Spin()
}
