package controllers

import (
	"WebApp/logic"
	"WebApp/models"
	"WebApp/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
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
	logic.PostVote()
	appG.Response(http.StatusOK,200,"")

}
