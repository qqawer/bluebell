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
	PostID    int64 `json:"post_id,string" binding:"required"`                //帖子id
	Direction int   `json:"direction,string" binding:"required,oneof=1 0 -1"` //赞成票(1)还是反对票(-1)取消投票(0)
}
