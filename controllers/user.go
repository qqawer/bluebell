package controllers

import (
	"WebApp/logic"
	"WebApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//SignUpHandler 处理注册请求的函数
func SignupHandler(c *gin.Context){
	//1.获取参数和参数
	// var user models.ParamSignUp
	p:=new(models.ParamSignUp)
	if err:=c.ShouldBindJSON(p);err!=nil{
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		c.JSON(http.StatusBadRequest,gin.H{"msg":"请求参数有误"})
		return
	}
	//2.业务处理
	if err:=logic.SignUp(p);err!=nil{
		c.JSON(http.StatusOK,gin.H{"msg":err.Error()})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK,gin.H{
		"msg":"Success",
	})
}