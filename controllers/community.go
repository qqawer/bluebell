package controllers

import (
	"WebApp/logic"
	"strconv"
	// "WebApp/models"
	"WebApp/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


// 跟社区相关的
func CommunityHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		// input struct {
		// 	ID   int64  `binding:"required" json:"id"`
		// 	Name string `binding:"required" json:"name"`
		// }
	)

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	zap.L().Error("Search with invalid param", zap.Error(err))
	// 	appG.Response(http.StatusOK, 400, "")
	// 	return
	// }
	// 	// 将输入转换为业务模型
	// 	community := models.Community{
	// 		ID:   input.ID,
	// 		Name: input.Name,
	// 	}
	//查询到所有的社区(community_id,community_name)以列表的形式的返回
	// data, err := logic.GetCommunityList(&community)
	data, err := logic.GetCommunityList()

	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		appG.Response(http.StatusOK, 500, "") // 不轻易把服务端报错暴露给外面
		return
	}

	appG.Response(http.StatusOK, 200, data)
}
func CommunityDetailHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	//1.获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err!=nil{
		appG.Response(http.StatusOK,400,"")
		return
	}
	//2.根据id获取社区详情
	data, err := logic.GetCommunitDetail(id)

	if err != nil {
		zap.L().Error("logic.GetCommunitDetail(id) failed", zap.Error(err))
		appG.Response(http.StatusOK, 500, "") // 不轻易把服务端报错暴露给外面
		return
	}

	appG.Response(http.StatusOK, 200, data)
}