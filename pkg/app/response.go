package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}

func (g *Gin) Response2(httpCode, errCode int, errMsg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  errMsg,
		Data: data,
	})
	return
}