package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

var (
	ErrorExitsEmailMsg = errors.New("该邮箱已存在")
	ErrorNotFindUser   = errors.New("登录失败，用户名或密码错误")
	ErrorCode          = "23505"
)

func NewUserDao(mdb *gorm.DB) *UserDao {
	return &UserDao{
		db: mdb,
	}
}

func (u *UserDao) InsertUser(ctx context.Context, user *User) error {
	now := time.Now().UnixMilli()
	user.CreatedAt = now
	user.UpdatedAt = now
	err := gorm.G[User](u.db).Create(ctx, user)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == ErrorCode {
			return fmt.Errorf("%w, 邮箱=%s", ErrorExitsEmailMsg, user.Email)
		}
		return errors.New("用户创建失败")
	}
	return err
}

func (u *UserDao) FindUserByEmail(ctx context.Context, email string) (User, error) {
	users, err := gorm.G[User](u.db).Where("email = ?", email).Find(ctx)
	if err != nil {
		return User{}, err
	}
	if len(users) == 0 {
		return User{}, fmt.Errorf("%w", ErrorNotFindUser)
	}
	return users[0], nil
}

// FindUserByID 根据用户 ID 查询用户
func (u *UserDao) FindUserByID(ctx context.Context, id int64) (User, error) {
	user, err := gorm.G[User](u.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, fmt.Errorf("%w", ErrorNotFindUser)
		}
		return User{}, err
	}
	return user, nil
}

// UpdatePassword 更新用户密码
func (u *UserDao) UpdatePassword(ctx context.Context, id int64, hashPassword string) error {
	_, err := gorm.G[User](u.db).Where("id = ?", id).Updates(ctx, User{
		Password:  hashPassword,
		UpdatedAt: time.Now().UnixMilli(),
	})
	return err
}
