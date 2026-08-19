package service

import (
	"context"
	"errors"
	"time"

	"github.com/codeLambQ/codeLamb_book/backend/internal/model"
	"github.com/codeLambQ/codeLamb_book/backend/internal/repository"
	"github.com/codeLambQ/codeLamb_book/backend/pkg/session"
)

// ErrorSessionExpired 会话已过期
var ErrorSessionExpired = errors.New("会话已过期")

type SessionService struct {
	Repo *repository.SessionRepository
	TTL  time.Duration // 会话有效期
}

func NewSessionService(repo *repository.SessionRepository) *SessionService {
	return &SessionService{
		Repo: repo,
		TTL:  7 * 24 * time.Hour,
	}
}

// Create 登录成功后创建会话，返回 sessionID
func (s *SessionService) Create(ctx context.Context, userID int64) (string, error) {
	sessionID, err := session.GenerateID()
	if err != nil {
		return "", err
	}

	sess := &model.Session{
		SessionID: sessionID,
		UserID:    userID,
		ExpiresAt: time.Now().Add(s.TTL).UnixMilli(),
	}
	if err := s.Repo.CreateSession(ctx, sess); err != nil {
		return "", err
	}
	return sessionID, nil
}

// Verify 校验会话是否有效，返回用户 ID
func (s *SessionService) Verify(ctx context.Context, sessionID string) (int64, error) {
	sess, err := s.Repo.FindSessionByID(ctx, sessionID)
	if err != nil {
		return 0, err
	}
	if sess.ExpiresAt < time.Now().UnixMilli() {
		return 0, ErrorSessionExpired
	}
	return sess.UserID, nil
}

// Logout 删除会话
func (s *SessionService) Logout(ctx context.Context, sessionID string) error {
	return s.Repo.DeleteSession(ctx, sessionID)
}
