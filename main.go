package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/common/util"
	"github.com/iprologue/myBlog/models"
	"github.com/iprologue/myBlog/pkg/setting"
	"github.com/iprologue/myBlog/router"
	"log"
	"syscall"
)

func init() {
	setting.SetUp()
	models.SetUp()
	util.SetUp()
}

func main() {

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	engine := gin.New()
	gin.SetMode(setting.ServerSetting.RunMode)
	router.InitRouter(engine)

	server := endless.NewServer(setting.ServerSetting.HttpPort, engine)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
