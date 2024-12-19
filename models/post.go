package models

import "time"

type Post struct {
	ID          int64      `json:"id,string" gorm:"column:post_id"`
	AuthorID    int64      `json:"author_id,string" gorm:"column:author_id"`
	CommunityID int64      `json:"community_id" gorm:"column:community_id" binding:"required"`
	Status      int32      `json:"status" gorm:"column:status"`
	Title       string     `json:"title" gorm:"column:title" binding:"required"`
	Content     string     `json:"content" gorm:"column:content" binding:"required"`
	CreateTime  *time.Time `json:"create_time" gorm:"column:create_time;autoUpdateTime"`
	UpdateTime  *time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

func (Post) TableName() string {
	return "post"
}

type ApiPostDetail struct {
	AuthorName string
	*Post      //嵌入帖子结构体
	*CommunityDetail  `json:"community"`  //嵌入社区信息
}
func (ApiPostDetail) TableName() string {
	return "post"
}
