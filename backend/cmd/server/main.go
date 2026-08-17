package main

import (
	"github.com/codeLambQ/codeLamb_book/backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	userHandler := handler.NerUserHandler()
	userHandler.RegisterUserHandler(server)

	err := server.Run()
	if err != nil {
		panic("启动失败: " + err.Error())
	}
}
