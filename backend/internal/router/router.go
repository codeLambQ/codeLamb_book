package router

import (
	"github.com/gin-gonic/gin"

	"github.com/codeLambQ/codeLamb_book/backend/internal/handler"
	"github.com/codeLambQ/codeLamb_book/backend/internal/middleware"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
	"github.com/codeLambQ/codeLamb_book/backend/internal/service"
)

// New 组装并返回 gin 引擎。
func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging())

	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)

	health := handler.NewHealthHandler()
	users := handler.NewUserHandler(svc)

	r.GET("/healthz", health.Check)

	api := r.Group("/api/v1")
	{
		api.GET("/users/:id", users.Get)
		api.POST("/users", users.Create)
	}

	return r
}
