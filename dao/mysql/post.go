package mysql

import "WebApp/models"

func CreatePost(p *models.Post)(error){
	if err:=db.Create(&p).Error;err!=nil{
		return err
	}
	return nil
}