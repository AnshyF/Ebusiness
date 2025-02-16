package router

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterRoutes(h *server.Hertz) {
	h.Use(
		middleware.RequestLogger(), // 请求日志（第一个执行）
		middleware.CORS(),          // 跨域处理
		middleware.Recovery(),      // 异常恢复（最后一个执行）
	)

}
