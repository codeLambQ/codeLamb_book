package middleware

import (
	"net/http"
	"time"

	"github.com/codeLambQ/codeLamb_book/backend/pkg/logger"
)

// Logging 请求日志中间件。
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info(r.Method, r.URL.Path, time.Since(start))
	})
}
