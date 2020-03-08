package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/pkg/setting"
	"github.com/iprologue/myBlog/router"
)

func main() {
	engine := gin.New()
	gin.SetMode(setting.ServerSetting.RunMode)
	router.InitRouter(engine)
	err := engine.Run(setting.ServerSetting.HttpPort)
	if err != nil {
		fmt.Println(err.Error())
	}

}
