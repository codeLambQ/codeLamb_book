package repository

import (
	"errors"
	"fmt"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository/dao"
	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	Dao *dao.UserDao
}

var ErrorExitsEmailMsg = dao.ErrorExitsEmailMsg

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		Dao: dao,
	}
}

func (u *UserRepository) RegisterUser(ctx *gin.Context, user *model.User) error {
	daoUser := &dao.User{Email: user.Email, Password: user.Password}
	err := u.Dao.InsertUser(ctx, daoUser)
	if errors.Is(err, ErrorExitsEmailMsg) {
		return fmt.Errorf("%w, 邮箱=%s", ErrorExitsEmailMsg, user.Email)
	}

	return err
}
