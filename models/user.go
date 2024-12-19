package models

type User struct {
	UserId   int64  `gorm:"column:user_id" json:"id,string"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password" json:"-"`
	Token string
}

func (User) TableName() string {
	return "user" // 映射到 user 表
}
