package logic

import (
	"WebApp/dao/mysql"
	"WebApp/dao/redis"
	"WebApp/models"
	"WebApp/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	//1.生成postID
	postID := snowflake.GenID()
	p.ID = postID

	//2，入库
	err:=mysql.CreatePost(p)
	if err!=nil{
		return err
	}
	return redis.CreatePost(p.ID)
	
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
		zap.L().Error("mysql.GetUserById(pots.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	//根据社区id拆线呢社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}
func GetPostList(page, size int) ([]*models.ApiPostDetail, error) {
	//获得list数据
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data := make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(pots.AuthorID) failed", zap.Int64("user_id", post.AuthorID), zap.Error(err))
			continue
		}
		//根据社区id拆线呢社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail:= &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data=append(data, postDetail)
	}
	return data,nil

}
