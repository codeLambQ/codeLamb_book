package repository

import (
	"context"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository/dao"
)

type SessionRepository struct {
	Dao *dao.SessionDao
}

func NewSessionRepository(d *dao.SessionDao) *SessionRepository {
	return &SessionRepository{Dao: d}
}

// CreateSession 创建会话
func (s *SessionRepository) CreateSession(ctx context.Context, sess *model.Session) error {
	daoSess := &dao.Session{
		SessionID: sess.SessionID,
		UserID:    sess.UserID,
		ExpiresAt: sess.ExpiresAt,
	}
	return s.Dao.InsertSession(ctx, daoSess)
}

// FindSessionByID 查询会话
func (s *SessionRepository) FindSessionByID(ctx context.Context, sessionID string) (*model.Session, error) {
	sess, err := s.Dao.FindSessionByID(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	return &model.Session{
		ID:        sess.ID,
		SessionID: sess.SessionID,
		UserID:    sess.UserID,
		ExpiresAt: sess.ExpiresAt,
		CreatedAt: sess.CreatedAt,
	}, nil
}

// DeleteSession 删除会话
func (s *SessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	return s.Dao.DeleteSession(ctx, sessionID)
}
