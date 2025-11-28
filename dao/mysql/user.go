package mysql

import (
	"WebApp/models"
	"WebApp/utils"
	"errors"
	"strings"

	"gorm.io/gorm"
)

//把每一步数据库封装成函数
//待logic层根据需求调用

// CheckUserExist 检查指定用户是否存在
func CheckUserExist(p *models.ParamSignUp) error {
    var user models.User
    err := db.Where("username=?", p.Username).Find(&user).Error // 使用 Find 避免报错

    if err != nil {
        if strings.Contains(err.Error(), "no such table") || strings.Contains(err.Error(), "doesn't exist") {
            return errors.New("数据库表不存在")
        }
        // 其他数据库错误直接返回
        return err
    }

    // 如果用户已存在
    if user.UserId != 0 { // 通过 UserId 判断是否找到用户
        return errors.New("用户已存在")
    }

    return nil // 用户不存在，允许注册
}


// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(input *models.User) (err error) {
	//对密码进行加密
	input.Password, err = utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	//执行sql语句入库
	if err := db.Create(input).Error; err != nil {
		return err
	}
	return nil
}
func Login(input *models.User)(error){
	// var user models.User
	oPassword:=input.Password
	//找不到返回返回
	if err := db.Where("username=?", input.Username).First(&input).Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return errors.New("用户不存在")
		}else{
			return err
		}
	}
	//验证密码
	if !utils.CheckPassword(oPassword,input.Password){
		return errors.New("密码错误")
	}
	return nil

}
func GetUserById(uid int64)(*models.User,error){
	var user models.User
	if err:=db.Where("user_id=?",uid).First(&user).Error;err!=nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,errors.New("找不到此用户")
		}else{
			return nil,err
		}
	}
	return &user, nil
}