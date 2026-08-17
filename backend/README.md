# backend

Go 后端微服务骨架（Gin 框架，标准目录结构）。

## 目录结构

    backend/
    ├── api/            # 接口文档（OpenAPI）
    ├── cmd/server/     # 程序入口
    ├── configs/        # 配置文件
    ├── internal/       # 私有代码（不对外导出）
    │   ├── config/     # 配置加载
    │   ├── handler/    # HTTP 处理层（gin handler）
    │   ├── middleware/ # 中间件（gin middleware）
    │   ├── model/      # 数据模型
    │   ├── repository/ # 数据访问层
    │   ├── router/     # 路由注册（gin engine）
    │   └── service/    # 业务逻辑层
    ├── migrations/     # 数据库迁移脚本
    ├── pkg/            # 可复用的公共包
    │   ├── logger/     # 日志
    │   └── response/   # 统一响应
    ├── scripts/        # 构建脚本
    ├── Dockerfile
    └── Makefile

## 快速开始

    make run        # 或 go run ./cmd/server
    curl http://localhost:8080/healthz

## 接口示例

    curl http://localhost:8080/api/v1/users/u1 -X POST -H "Content-Type: application/json" -d '{"id":"u1","name":"alice","email":"a@b.com"}'
    curl http://localhost:8080/api/v1/users/u1

## 常用命令

    make build      # 构建二进制到 bin/
    make test       # 运行测试
    make tidy       # 整理依赖

