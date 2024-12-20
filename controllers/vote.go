package controllers

import (
	"WebApp/logic"
	"WebApp/models"
	"WebApp/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostVoteController(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		appG.Response2(http.StatusOK,400,err.Error(),"")
		return
	}
	// 获取当前请求的用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		appG.Response2(http.StatusOK,500,"请先登录用户","")
		return
	}
	// 具体投票的业务逻辑
	if err:=logic.PostVote(userID,p);err!=nil{
		zap.L().Error("ogic.PostVote(userID,p) failed",zap.Error(err))
		appG.Response2(http.StatusOK,500,err.Error(),"")
		return
	}
	appG.Response(http.StatusOK,200,"")

}
