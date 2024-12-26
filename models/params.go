package models

//定义请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username" bingding:"required"`
	Password   string `json:"password" bingding:"required"`
	RePassword string `json:"re_password" bingding:"required"`
}
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`                //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票(1)还是反对票(-1)取消投票(0)
}


type ParamPostList struct{
	Page int64 `form:"page"`
	Size int64	`form:"size"`
	Order string `form:"order"`
}

type ParamCommunityPostList struct{
	ParamPostList
	CommunityID int64 `form:"community_id"`
}