package logic

import (
	"WebApp/dao/mysql"
	"WebApp/models"
	"WebApp/pkg/snowflake"
)

// 存放业务逻辑代码
func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户在不在

	if err := mysql.CheckUserExist(p);err!=nil{
	//数据库查询出错
	return err
	}
	
	//生成Uid
	userID := snowflake.GenID()
	input:=&models.User{
		UserId: userID, 
		Username :p.Username,
		Password :p.Password,
	}
	

	//保存进数据库(密码加密)
	return mysql.InsertUser(input)

}
