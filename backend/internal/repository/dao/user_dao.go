package dao

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

var (
	ErrorExitsEmailMsg = errors.New("该邮箱已存在")
	ErrorCode          = "23505"
)

func NewUserDao(mdb *gorm.DB) *UserDao {
	return &UserDao{
		db: mdb,
	}
}

func (u *UserDao) InsertUser(ctx *gin.Context, user *User) error {
	err := gorm.G[User](u.db).Create(ctx, user)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == ErrorCode {
			return fmt.Errorf("%w, 邮箱=%s", ErrorExitsEmailMsg, user.Email)
		}
		return errors.New("用户创建失败")
	}
	return err
}
