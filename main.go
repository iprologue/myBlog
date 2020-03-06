package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/router"
)

func main() {
	engine := gin.New()
	router.InitRouter(engine)
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}

}
