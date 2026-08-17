package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/codeLambQ/codeLamb_book/backend/pkg/logger"
)

// Logging 请求日志中间件。
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info(c.Request.Method, c.FullPath(), c.Writer.Status(), time.Since(start))
	}
}
