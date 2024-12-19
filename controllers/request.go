package controllers

import (
	"WebApp/middlewares"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getCurrentUserID 获取当前登录的用户ID
func getCurrentUserID(c *gin.Context) (UserID int64, err error) {
	uid, ok := c.Get(middlewares.CtxUserIDKey)
	if !ok {
		err = errors.New("用户未登录")
		return
	}
	userID, ok := uid.(int64)
	if !ok {
		err = errors.New("用户ID类型断言失败")
		return
	}
	return userID,nil
}

func getPageInfo(c *gin.Context)(int,int){
	pageStr:=c.Query("page")
	sizeStr:=c.Query("size")
	
	page,err:=strconv.Atoi(pageStr)
	if err!=nil{
		page=1
	}
	size,err:=strconv.Atoi(sizeStr)
	if err!=nil{
		size=10
	}
	return page,size

}