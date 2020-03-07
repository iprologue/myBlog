package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/pkg/setting"
	"github.com/iprologue/myBlog/router"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(setting.RunMode)
	router.InitRouter(engine)
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}

}
