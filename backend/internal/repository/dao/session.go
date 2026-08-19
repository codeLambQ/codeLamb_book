package dao

// Session 会话入库结构体
type Session struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	SessionID string `gorm:"uniqueIndex;size:64"`
	UserID    int64  `gorm:"index"`
	ExpiresAt int64  // 过期时间戳（毫秒）
	CreatedAt int64
}
