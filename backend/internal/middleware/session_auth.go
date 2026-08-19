package middleware

import (
	"net/http"

	"github.com/codeLambQ/codeLamb_book/backend/internal/service"
	"github.com/gin-gonic/gin"
)

// ctxUserIDKey 上下文中存储当前登录用户 ID 的 key
const ctxUserIDKey = "user_id"

// SessionAuth 登录认证中间件：解析 Cookie 中的 session_id，校验并注入 user_id
func SessionAuth(sessionSvc *service.SessionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionID, err := ctx.Cookie("session_id")
		if err != nil || sessionID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "未登录"})
			return
		}

		userID, err := sessionSvc.Verify(ctx, sessionID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "登录已过期，请重新登录"})
			return
		}

		ctx.Set(ctxUserIDKey, userID)
		ctx.Next()
	}
}

// GetUserID 从上下文中取出当前登录用户 ID
func GetUserID(ctx *gin.Context) (int64, bool) {
	v, ok := ctx.Get(ctxUserIDKey)
	if !ok {
		return 0, false
	}
	userID, ok := v.(int64)
	return userID, ok
}
