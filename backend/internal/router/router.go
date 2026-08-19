package router

import (
	"github.com/codeLambQ/codeLamb_book/backend/internal/handler"
	"github.com/codeLambQ/codeLamb_book/backend/internal/middleware"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository/dao"
	"github.com/codeLambQ/codeLamb_book/backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	server := gin.Default()
	server.Use(middleware.GlobalCors())

	// 用户模块依赖
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	userService := service.NewUserService(ur)

	// 会话模块依赖
	sd := dao.NewSessionDao(db)
	sr := repository.NewSessionRepository(sd)
	sessionService := service.NewSessionService(sr)

	userHandler := handler.NewUserHandler(userService, sessionService)
	userHandler.RegisterUserHandler(server)

	// 需要登录的接口
	auth := server.Group("/")
	auth.Use(middleware.SessionAuth(sessionService))
	{
		auth.GET("/me", userHandler.Me)
		auth.GET("/users/:id", userHandler.Profile)
		auth.POST("/users/:id", userHandler.Edit)
	}

	return server
}
