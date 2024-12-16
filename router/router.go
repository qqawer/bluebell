package router

import (
	"WebApp/controllers"
	"WebApp/logger"
	"WebApp/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 使用跨域中间件
	// r.Use(middlewares.CORSMiddleware())
	v1 := r.Group("/api/v1")

	v1.POST("/signup", controllers.SignupHandler)
	v1.GET("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id",controllers.CommunityDetailHandler)
	}

	r.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"msg":"404",
		})
	})

	return r
}
