package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/service"
	"github.com/codeLambQ/codeLamb_book/backend/pkg/response"
)

// UserHandler 用户 HTTP 处理层。
type UserHandler struct {
	svc *service.UserService
}

// NewUserHandler 创建处理器。
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// Get 查询用户：GET /api/v1/users/:id
func (h *UserHandler) Get(c *gin.Context) {
	u, err := h.svc.Get(c.Param("id"))
	if err != nil {
		response.JSON(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	response.JSON(c, http.StatusOK, "ok", u)
}

// Create 创建用户：POST /api/v1/users
func (h *UserHandler) Create(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		response.JSON(c, http.StatusBadRequest, "invalid body", nil)
		return
	}
	if u.ID == "" {
		response.JSON(c, http.StatusBadRequest, "id is required", nil)
		return
	}
	_ = h.svc.Create(u)
	response.JSON(c, http.StatusCreated, "created", nil)
}
