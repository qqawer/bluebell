package models

// import "time"

// type Community struct {
//     ID           int64     `json:"id" db:"id"`
//     CommunityID  int64     `json:"community_id" db:"community_id"`
//     Name         string    `json:"name" db:"community_name"`
//     Introduction string    `json:"introduction" db:"introduction"`
//     CreateTime   time.Time `json:"create_time" db:"create_time"`
//     UpdateTime   time.Time `json:"update_time" db:"update_time"`
// }

type Community struct {
    ID           int64  `json:"community_id" gorm:"column:community_id"`
    Name         string `json:"name" gorm:"column:community_name"`
}

func (Community) TableName() string {
	return "community" // 映射到 user 表
}


