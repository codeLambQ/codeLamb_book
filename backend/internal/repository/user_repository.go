package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository/dao"
)

type UserRepository struct {
	Dao *dao.UserDao
}

var ErrorExitsEmailMsg = dao.ErrorExitsEmailMsg
var ErrorNotFindUser = dao.ErrorNotFindUser

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		Dao: dao,
	}
}

func (u *UserRepository) RegisterUser(ctx context.Context, user *model.User) error {
	daoUser := &dao.User{Email: user.Email, Password: user.Password}
	err := u.Dao.InsertUser(ctx, daoUser)
	if errors.Is(err, ErrorExitsEmailMsg) {
		return fmt.Errorf("%w, 邮箱=%s", ErrorExitsEmailMsg, user.Email)
	}

	return err
}

func (u *UserRepository) FindUserForEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.Dao.FindUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}
	if user.Email == "" {
		return nil, fmt.Errorf("%w", ErrorNotFindUser)
	}
	return &model.User{ID: user.ID, Email: user.Email, Password: user.Password}, nil
}

// FindUserByID 根据用户 ID 查询用户
func (u *UserRepository) FindUserByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.Dao.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.User{ID: user.ID, Email: user.Email, Password: user.Password}, nil
}

// UpdatePassword 更新用户密码
func (u *UserRepository) UpdatePassword(ctx context.Context, id int64, hashPassword string) error {
	return u.Dao.UpdatePassword(ctx, id, hashPassword)
}
