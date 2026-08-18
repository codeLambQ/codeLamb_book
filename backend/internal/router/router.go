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
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	userService := service.NewUserService(ur)
	userHandler := handler.NewUserHandler(userService)
	userHandler.RegisterUserHandler(server)

	return server
}
