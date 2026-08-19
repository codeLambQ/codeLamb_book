package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ErrorSessionNotFound 会话不存在
var ErrorSessionNotFound = errors.New("会话不存在")

type SessionDao struct {
	db *gorm.DB
}

func NewSessionDao(mdb *gorm.DB) *SessionDao {
	return &SessionDao{db: mdb}
}

// InsertSession 创建会话
func (s *SessionDao) InsertSession(ctx context.Context, sess *Session) error {
	sess.CreatedAt = time.Now().UnixMilli()
	return gorm.G[Session](s.db).Create(ctx, sess)
}

// FindSessionByID 根据 sessionID 查询会话
func (s *SessionDao) FindSessionByID(ctx context.Context, sessionID string) (Session, error) {
	sessions, err := gorm.G[Session](s.db).Where("session_id = ?", sessionID).Find(ctx)
	if err != nil {
		return Session{}, err
	}
	if len(sessions) == 0 {
		return Session{}, fmt.Errorf("%w", ErrorSessionNotFound)
	}
	return sessions[0], nil
}

// DeleteSession 删除会话
func (s *SessionDao) DeleteSession(ctx context.Context, sessionID string) error {
	_, err := gorm.G[Session](s.db).Where("session_id = ?", sessionID).Delete(ctx)
	return err
}
