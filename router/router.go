package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/service"
)

func InitRouter(r *gin.Engine)  {

	GroupV1 := r.Group("/api/v1")
	{
		GroupV1.GET("/tags", v1.GetTags)
		GroupV1.POST("/tags", v1.AddTag)
		GroupV1.PUT("/tags/:id", v1.EditTag)
		GroupV1.DELETE("/tags/:id", v1.DeletedTag)
	}
}
