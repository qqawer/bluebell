package logic

import (
	"WebApp/dao/mysql"
	"WebApp/models"
	"WebApp/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	//1.生成postID
	postID := snowflake.GenID()
	p.ID = postID

	//2，入库
	return mysql.CreatePost(p)
}
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	data = new(models.ApiPostDetail)
	
	//查询并组合我们接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	//根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(pots.AuthorID) failed", zap.Error(err))
		return
	}
	//根据社区id拆线呢社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}
	data=&models.ApiPostDetail{
		AuthorName: user.Username,
		Post: post,
		CommunityDetail: community,
	}
	return
}
