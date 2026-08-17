package service

import (
	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
)

// UserService 用户业务逻辑层。
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService 创建服务。
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Get 查询用户。
func (s *UserService) Get(id string) (model.User, error) {
	return s.repo.Get(id)
}

// Create 创建用户。
func (s *UserService) Create(u model.User) error {
	s.repo.Save(u)
	return nil
}
