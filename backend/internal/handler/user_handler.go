package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/codeLambQ/codeLamb_book/backend/internal/middleware"
	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/service"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

const (
	EmailRegexp    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	PasswordRegexp = `^(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]).{8,}$`
)

type UserHandler struct {
	Svc            *service.UserService
	SessionSvc     *service.SessionService
	EmailRegexp    *regexp2.Regexp
	PasswordRegexp *regexp2.Regexp
}

func NewUserHandler(svc *service.UserService, sessionSvc *service.SessionService) *UserHandler {
	emailRegexp := regexp2.MustCompile(EmailRegexp, regexp2.None)
	passwordRegexp := regexp2.MustCompile(PasswordRegexp, regexp2.None)

	return &UserHandler{
		Svc:            svc,
		SessionSvc:     sessionSvc,
		EmailRegexp:    emailRegexp,
		PasswordRegexp: passwordRegexp,
	}
}

// RegisterUserHandler 注册公开路由（无需登录）
func (u *UserHandler) RegisterUserHandler(server *gin.Engine) {
	server.POST("/users", u.Register)
	server.POST("/login", u.Login)
	server.POST("/logout", u.Logout)
}

func (u *UserHandler) Register(ctx *gin.Context) {
	req := &struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}

	// 获取 restful 请求的请求体内容，传输到内部 struct req 上
	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "系统错误，请稍后重试",
		})
		return
	}

	// 做 email password 校验
	ok, err := u.EmailRegexp.MatchString(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统错误，请稍后重试",
		})
		return
	}

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "邮件格式有错误，请重新输入",
		})
		return
	}

	// 密码一致性判断
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致！",
		})
		return
	}
	ok, err = u.PasswordRegexp.MatchString(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统错误，请稍后重试",
		})
		return
	}

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "密码格式有误：长度至少为8位，必须包含大写字母、小写字母和特殊字符",
		})
		return
	}

	err = u.Svc.RegisterUser(ctx, &model.User{Email: req.Email, Password: req.Password})
	if err != nil {
		if errors.Is(err, service.ErrorExitsEmailMsg) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "邮箱已存在",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "创建用户失败",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "注册成功，即将跳转到登录界面",
	})
}

// Login 用户登录，成功后创建会话并下发 Cookie
func (u *UserHandler) Login(ctx *gin.Context) {
	req := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "系统错误，请稍后重试",
		})
		return
	}

	user, err := u.Svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrorNotFindUser) || errors.Is(err, service.ErrorPasswordNotAccor) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "登录失败，邮箱或密码不正确"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误，请稍后重试"})
		return
	}

	// 登录成功：创建会话
	sessionID, err := u.SessionSvc.Create(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误，请稍后重试"})
		return
	}

	// 下发 Cookie：HttpOnly 防 XSS；SameSite 防 CSRF；生产环境 secure 需改为 true
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("session_id", sessionID, int(u.SessionSvc.TTL.Seconds()), "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"email":   user.Email,
	})
}

// Logout 退出登录：删除服务端会话并清除 Cookie
func (u *UserHandler) Logout(ctx *gin.Context) {
	if sessionID, err := ctx.Cookie("session_id"); err == nil && sessionID != "" {
		_ = u.SessionSvc.Logout(ctx, sessionID)
	}
	ctx.SetCookie("session_id", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "已退出登录"})
}

// Profile 查看个人信息
func (u *UserHandler) Profile(ctx *gin.Context) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "未登录"})
		return
	}

	// 只能查看自己的信息
	paramID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || paramID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "无权访问"})
		return
	}

	user, err := u.Svc.GetProfile(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "用户不存在"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

// Edit 修改个人信息（当前支持修改密码）
func (u *UserHandler) Edit(ctx *gin.Context) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "未登录"})
		return
	}

	paramID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || paramID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "无权访问"})
		return
	}

	req := &struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}{}
	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "系统错误，请稍后重试"})
		return
	}

	ok, err = u.PasswordRegexp.MatchString(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误，请稍后重试"})
		return
	}
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "密码格式有误：长度至少为8位，必须包含大写字母、小写字母和特殊字符"})
		return
	}

	if err := u.Svc.ChangePassword(ctx, userID, req.OldPassword, req.NewPassword); err != nil {
		if errors.Is(err, service.ErrorPasswordNotAccor) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "旧密码不正确"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "修改密码失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})
}
