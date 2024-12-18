package logic

import (
	"WebApp/dao/mysql"
	"WebApp/models"
	"WebApp/pkg/snowflake"
)

func CreatePost( p *models.Post)(error){
	//1.生成postID
	postID:=snowflake.GenID()
	p.ID=postID

	//2，入库
	return mysql.CreatePost(p)
}
func GetDetailPost(pid int64)(data *models.Post,err error){
	return mysql.GetDatailPost(pid)
}