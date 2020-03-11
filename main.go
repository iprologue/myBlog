package main

import (
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/iprologue/myBlog/common/util"
	"github.com/iprologue/myBlog/models"
	"github.com/iprologue/myBlog/pkg/gredis"
	"github.com/iprologue/myBlog/pkg/setting"
	"github.com/iprologue/myBlog/router"
)

func init() {
	setting.SetUp()
	models.SetUp()
	gredis.SetUp()
	util.SetUp()
}

func main() {

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	engine := router.InitRouter()
	server := endless.NewServer(setting.ServerSetting.HttpPort, engine)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
