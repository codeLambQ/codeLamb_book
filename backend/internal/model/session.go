package model

// Session 会话领域模型
type Session struct {
	ID        int64  `json:"id"`
	SessionID string `json:"session_id"`
	UserID    int64  `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"` // 过期时间戳（毫秒）
	CreatedAt int64  `json:"created_at"` // 创建时间戳（毫秒）
}
