package repository

import (
	"errors"
	"sync"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
)

// ErrNotFound 记录不存在。
var ErrNotFound = errors.New("user not found")

// UserRepository 用户数据访问层（内存实现，可替换为数据库）。
type UserRepository struct {
	mu    sync.RWMutex
	users map[string]model.User
}

// NewUserRepository 创建仓库。
func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]model.User)}
}

// Get 按 ID 查询用户。
func (r *UserRepository) Get(id string) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return model.User{}, ErrNotFound
	}
	return u, nil
}

// Save 保存用户。
func (r *UserRepository) Save(u model.User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[u.ID] = u
}
