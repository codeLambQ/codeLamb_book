package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Body 统一响应体。
type Body struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// OK 成功响应。
func OK(c *gin.Context, data any) {
	JSON(c, http.StatusOK, "ok", data)
}

// JSON 写入统一响应。
func JSON(c *gin.Context, status int, msg string, data any) {
	c.JSON(status, Body{Code: status, Msg: msg, Data: data})
}
