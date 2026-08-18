package service

import (
	"errors"
	"fmt"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var ErrorExitsEmailMsg = repository.ErrorExitsEmailMsg

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: ur,
	}
}

// RegisterUser 注册用户逻辑
func (u *UserService) RegisterUser(ctx *gin.Context, user *model.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPassword)
	err = u.UserRepository.RegisterUser(ctx, user)
	if errors.Is(err, ErrorExitsEmailMsg) {
		return fmt.Errorf("%w, 邮箱=%s", ErrorExitsEmailMsg, user.Email)
	}

	return err
}
