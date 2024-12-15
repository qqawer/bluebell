package controllers

import (
	"WebApp/logic"
	"WebApp/models"
	"WebApp/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignupHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	//1.获取参数和参数
	// var user models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// c.JSON(http.StatusBadRequest,gin.H{"msg":"请求参数有误"})
		appG.Response(http.StatusOK, 400, "")
		return
	}
	//2.业务处理
	if err := logic.SignUp(p); err != nil {

		appG.Response2(http.StatusOK, 400, err.Error(), "")
		zap.L().Error("Logic.Signup failed", zap.Error(err))
		return
	}
	//3.返回响应
	appG.Response(http.StatusOK, 200, "")
}
func LoginHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	//1.验证表单
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// c.JSON(http.StatusBadRequest,gin.H{"msg":"请求参数有误"})
		appG.Response(http.StatusOK, 400, "")
		return
	}
	//2.逻辑业务
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed ", zap.String("username", p.Username), zap.Error(err))
		// c.JSON(http.StatusOK,gin.H{"msg":"用户名或密码错误"})
		appG.Response2(http.StatusOK, 400, err.Error(), "")
		return
	}
	//3.返回数据
	// c.JSON(http.StatusOK,gin.H{"msg":"登陆成功"})
	appG.Response(http.StatusOK, 200, token)

}
