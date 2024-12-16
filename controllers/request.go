package controllers

import (
	"WebApp/middlewares"
	"errors"

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
