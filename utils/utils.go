// utils/utils.go
package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// SendError 向客户端发送错误响应
func SendError(c *app.RequestContext, statusCode int, message string) {
	resp := map[string]interface{}{
		"code":    statusCode,
		"message": message,
	}
	c.JSON(statusCode, resp)
}

// SendSuccess 向客户端发送成功响应
func SendSuccess(c *app.RequestContext, message string, data interface{}) {
	resp := map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
	c.JSON(http.StatusOK, resp)
}
