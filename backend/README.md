# backend

Go 后端微服务（Gin + GORM + PostgreSQL）。

## 技术栈

- Gin v1.12
- GORM v1.31 + PostgreSQL
- golang.org/x/crypto（bcrypt 密码哈希）
- dlclark/regexp2（正则校验）

## 目录结构

    backend/
    ├── cmd/server/          # 程序入口
    ├── internal/
    │   ├── handler/         # HTTP 处理器
    │   ├── service/         # 业务逻辑
    │   ├── repository/      # 仓储（含 dao）
    │   ├── model/           # 领域模型
    │   ├── middleware/      # CORS、会话认证等中间件
    │   └── router/          # 路由组装
    └── pkg/                 # 通用工具

## 运行

1. 准备 PostgreSQL，修改 `cmd/server/main.go` 中的 DSN。
2. 运行 `go run ./cmd/server`，启动后自动建表（users、sessions）。

## 接口

| 方法 | 路径          | 是否登录 | 说明                         |
|------|---------------|----------|------------------------------|
| POST | /users        | 否       | 注册                         |
| POST | /login        | 否       | 登录，成功后下发 session Cookie |
| POST | /logout       | 否       | 退出登录                     |
| GET  | /users/:id    | 是       | 查看个人信息                 |
| POST | /users/:id    | 是       | 修改个人信息（修改密码）     |

## 会话方案（Cookie + Session）

- 登录成功：服务端生成随机 `sessionID` 存入 `sessions` 表，通过 `Set-Cookie` 下发。
- 后续请求：浏览器自动携带 Cookie，`SessionAuth` 中间件校验会话并注入 `user_id`。
- Cookie 属性：`HttpOnly`、`SameSite=Lax`；生产环境建议开启 `Secure`。
- 会话有效期默认 7 天，在 `service.NewSessionService` 中配置。
