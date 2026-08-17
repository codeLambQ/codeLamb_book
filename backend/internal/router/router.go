package router

import (
	"net/http"

	"github.com/codeLambQ/codeLamb_book/backend/internal/handler"
	"github.com/codeLambQ/codeLamb_book/backend/internal/middleware"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
	"github.com/codeLambQ/codeLamb_book/backend/internal/service"
)

// New 组装并返回路由。
func New() http.Handler {
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)

	health := handler.NewHealthHandler()
	users := handler.NewUserHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", health.Check)
	mux.HandleFunc("GET /api/v1/users/{id}", users.Get)
	mux.HandleFunc("POST /api/v1/users", users.Create)

	return middleware.Logging(mux)
}
