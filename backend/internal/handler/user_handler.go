package handler

import (
	"errors"
	"net/http"

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
	EmailRegexp    *regexp2.Regexp
	PasswordRegexp *regexp2.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	emailRegexp := regexp2.MustCompile(EmailRegexp, regexp2.None)
	passwordRegexp := regexp2.MustCompile(PasswordRegexp, regexp2.None)

	return &UserHandler{
		Svc:            svc,
		EmailRegexp:    emailRegexp,
		PasswordRegexp: passwordRegexp,
	}
}

func (u *UserHandler) RegisterUserHandler(server *gin.Engine) {
	server.POST("/users", u.Register)
	server.POST("/login", u.Login)
	server.POST("/users/:id", u.Edit)
	server.GET("/users/:id", u.Profile)
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

// Login 用户登录
func (u *UserHandler) Login(ctx *gin.Context) {
	// 定义一个内部结构体，获取用户请求的 body
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"email":   user.Email,
	})
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "功能未实现"})
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "功能未实现"})
}
