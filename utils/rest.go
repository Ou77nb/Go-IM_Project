package utils

import (
	"github.com/gin-gonic/gin"
)

// Response 统一返回结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 返回信息
	Data    interface{} `json:"data"`    // 返回数据
}

// Success 成功返回
func Success(c *gin.Context, code int, message string, data interface{}) {
	resp := Response{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
	if code != 0 {
		resp.Code = code
	}
	if len(message) > 0 {
		resp.Message = message
	}
	c.JSON(200, resp)
}

// Error 失败返回
func Error(c *gin.Context, code int, message string) {
	resp := Response{
		Code:    201,
		Message: message,
		Data:    nil,
	}
	if code != 0 {
		resp.Code = code
	}
	if len(message) > 0 {
		resp.Message = message
	}
	c.JSON(200, resp)
}
