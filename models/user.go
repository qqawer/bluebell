package models

type User struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	UserId   int64  `gorm:"column:user_id"`
}

func (User) TableName() string {
	return "user" // 映射到 user 表
}
