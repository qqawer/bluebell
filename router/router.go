package router

import (
	"WebApp/controllers"
	"WebApp/logger"
	

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode==gin.ReleaseMode{
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 使用跨域中间件
	// r.Use(middlewares.CORSMiddleware())

	r.GET("/signup",controllers.SignupHandler)

	return r
}
