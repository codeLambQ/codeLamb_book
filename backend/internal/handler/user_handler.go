package handler

import (
	"fmt"
	"net/http"

	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

const (
	Email_Regexp    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	Password_Regexp = `^(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]).{8,}$`
)

type UserHandler struct {
	EmailRegexp    *regexp2.Regexp
	PasswordRegexp *regexp2.Regexp
}

func NerUserHandler() *UserHandler {
	emailRegexp := regexp2.MustCompile(Email_Regexp, regexp2.Debug)
	passwordRegexp := regexp2.MustCompile(Password_Regexp, regexp2.Debug)

	return &UserHandler{
		EmailRegexp:    emailRegexp,
		PasswordRegexp: passwordRegexp,
	}
}

func (u *UserHandler) RegisterUserHandler(server *gin.Engine) {
	server.PUT("/users", u.LogUp)
	server.POST("/users", u.Login)
	server.POST("/users/:id", u.Edit)
	server.DELETE("/users/:id", u.Profile)
}

func (u *UserHandler) LogUp(ctx *gin.Context) {
	req := &struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}

	// 获取 restful 请求的请求体内容，传输到内部 struct req 上
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": fmt.Sprintln("系统错误，请稍后重试"),
		})
		return
	}

	// 做 email password 校验
	ok, err := u.EmailRegexp.MatchString(req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": fmt.Sprintln("系统错误，请稍后重试"),
		})
		return
	}

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintln("邮件格式有错误，请重新输入"),
		})
		return
	}

	// 密码一致性判断
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintln("两次密码不一致！"),
		})
		return
	}
	ok, err = u.PasswordRegexp.MatchString(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": fmt.Sprintln("系统错误，请稍后重试"),
		})
		return
	}

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintln("密码格式有误：长度最少为8位，必须包含大小写字符、特殊字符、数字"),
		})
		return
	}

	// TODO 注册操作

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintln("注册成功，即将跳转到登录界面"),
	})
	fmt.Printf("req body %+v", req)

}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
