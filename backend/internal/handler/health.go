package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler 健康检查。
type HealthHandler struct{}

// NewHealthHandler 创建处理器。
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check 健康检查接口。
func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
