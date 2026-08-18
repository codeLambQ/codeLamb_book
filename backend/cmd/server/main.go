package main

import (
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository/dao"
	"github.com/codeLambQ/codeLamb_book/backend/internal/router"
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

	server := router.NewRouter(db)

	err = server.Run()
	if err != nil {
		panic("启动失败: " + err.Error())
	}
}
