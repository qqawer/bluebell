package controllers

import (
	"WebApp/logic"
	"WebApp/models"
	"WebApp/pkg/app"
	"net/http"
	"strconv"

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
		zap.L().Debug("c.ShouldBindJSON(&post) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		appG.Response(http.StatusOK, 400, "")
		return
	}
	//2.获取当前登录用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		appG.Response2(http.StatusOK, 500, "You need login", "")
		zap.L().Error("getCurrentUserID(c) failed", zap.Error(err))
		return
	}
	p.AuthorID = userID
	//3.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		appG.Response(http.StatusOK, 500, "")
		return
	}
	//4.返回响应
	appG.Response(http.StatusOK, 200, "")
}
func GetDetailPostHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	//1，获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		appG.Response(http.StatusOK, 400, "")
		return
	}

	//2.获取数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetDetailPost(pid) failed", zap.Error(err))
		appG.Response(http.StatusOK, 500, "")
		return
	}
	//3.返回响应
	appG.Response(http.StatusOK, 200, data)
}
func GetPostListHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	//1.获取分页数据
	page, size := getPageInfo(c)
	//2.获取List数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList(page,size) failed", zap.Error(err))
		appG.Response(http.StatusOK, 500, "")
		return
	}
	//3.返回响应
	appG.Response(http.StatusOK, 200, data)
}

// 按时间、分数排序
//1.获取参数
//2.redis获取id列表
//3.根据id去获取获取list数据库
func GetPostListHandler2(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	//检查param
	var p *models.ParamPostList
	if err := c.ShouldBindQuery(&p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params",zap.Error(err))
		appG.Response(http.StatusOK,400,"")
	}
	
	data,err:=logic.GetPostList2(p)
	if err!=nil{
		zap.L().Error("logic.GetPostList2(p) failed",zap.Error(err))
		appG.Response(http.StatusOK,500,"")
	}
	//
	//返回响应
	appG.Response(http.StatusOK,200,data)
}

//根据社区去查询帖子列表
func GetCommunityPostListHandler(c *gin.Context){
	var (
		appG = app.Gin{C: c}
	)
	//检查param
	var p *models.ParamCommunityPostList
	if err := c.ShouldBindQuery(&p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid params",zap.Error(err))
		appG.Response(http.StatusOK,400,"")
	}
	
	data,err:=logic.GetCommunityPostList(p)
	if err!=nil{
		zap.L().Error("logic.GetPostList2(p) failed",zap.Error(err))
		appG.Response(http.StatusOK,500,"")
	}
	//
	//返回响应
	appG.Response(http.StatusOK,200,data)
}