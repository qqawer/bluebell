package mysql

import (
	"WebApp/models"
	"errors"

	"gorm.io/gorm"
)

func CreatePost(p *models.Post)(error){
	if err:=db.Create(&p).Error;err!=nil{
		return err
	}
	return nil
}
func GetDatailPost(pid int64)(*models.Post, error){
	var post models.Post
	if err:=db.Where("post_id=?",pid).First(&post).Error;err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,errors.New("cannot find data with this pid")
		}else{
			return nil,err
		}
	}
	return &post,nil
}