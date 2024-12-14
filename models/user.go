package models

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	UserId   int64    `db:"user_id"`
}
func (User) TableName() string {
    return "user"  // 映射到 user 表
}