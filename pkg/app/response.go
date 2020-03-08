package app

import (
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/common/errcode"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{}  `json:"data"`
}

func (g *Gin) Response(httpCode, errerCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:errerCode,
		Msg:errcode.GetMsg(errerCode),
		Data:data,
	})
	return
}