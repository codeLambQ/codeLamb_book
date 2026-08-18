package main

import (
	"github.com/codeLambQ/codeLamb_book/backend/internal/handler"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository/dao"
	"github.com/codeLambQ/codeLamb_book/backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	dsn := "host=localhost user=webook password=Psbc@970726.. dbname=webook port=15432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	err = dao.InitTable(db)
	if err != nil {
		panic("自动建表失败" + err.Error())
	}
	server := gin.Default()

	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	userService := service.NewUserService(ur)
	userHandler := handler.NewUserHandler(userService)
	userHandler.RegisterUserHandler(server)

	err = server.Run()
	if err != nil {
		panic("启动失败: " + err.Error())
	}
}
