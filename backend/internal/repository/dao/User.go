package dao

// User 入库数据库的结构体
type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	CreateAt int64
	UpdateAt int64
}
