package models

//定义请求参数的结构体
type ParamSignUp struct{
	Username string `json:"username" bingding:"required"`
	Password string `json:"password" bingding:"required"`
	RePassword string `json:"re_password" bingding:"required"`
}
type ParamLogin struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}