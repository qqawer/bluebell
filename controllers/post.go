package controllers

import (
	"WebApp/logic"
	"WebApp/models"
	"WebApp/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	//1.验证表单
	var p *models.Post
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(&post) error",zap.Any("err",err))
		zap.L().Error("create post with invalid param")
		appG.Response(http.StatusOK,400,"")
		return
	}
	//2.获取当前登录用户的ID
	userID,err:=getCurrentUserID(c)
	if err!=nil{
		appG.Response2(http.StatusOK,500,"You need login","")
		zap.L().Error("getCurrentUserID(c) failed", zap.Error(err))
		return
	}
	p.AuthorID=userID
	
	//3.创建帖子
	if err:=logic.CreatePost(p);err!=nil{
		zap.L().Error("logic.CreatePost(p) failed",zap.Error(err))
		appG.Response(http.StatusOK,500,"")
		return
	}
	//4.返回响应
	appG.Response(http.StatusOK,200,"")
}
