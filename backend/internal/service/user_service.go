package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrorExitsEmailMsg = repository.ErrorExitsEmailMsg
var ErrorNotFindUser = repository.ErrorNotFindUser
var ErrorPasswordNotAccor = errors.New("密码或用户名不正确")

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: ur,
	}
}

// RegisterUser 注册用户逻辑
func (u *UserService) RegisterUser(ctx context.Context, user *model.User) error {

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

// Login 用户登录业务
func (u *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := u.UserRepository.FindUserForEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("%w", ErrorNotFindUser)
	}

	// 判断密码是否一致
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, fmt.Errorf("%w", ErrorPasswordNotAccor)
	}
	return user, nil
}
