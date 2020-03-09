package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/iprologue/myBlog/docs"
	"github.com/iprologue/myBlog/middleware/jwt"
	"github.com/iprologue/myBlog/middleware/logger"
	"github.com/iprologue/myBlog/pkg/upload"
	"github.com/iprologue/myBlog/router/api"
	"github.com/iprologue/myBlog/router/api/v1"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter(r *gin.Engine) {

	r.Use(gin.Logger(), gin.Recovery(), logger.LoggerToFile())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	GroupV1 := r.Group("/api/v1").Use(jwt.JWT())
	{
		GroupV1.GET("/tags", v1.GetTags)
		GroupV1.POST("/tags", v1.AddTag)
		GroupV1.PUT("/tags/:id", v1.EditTag)
		GroupV1.DELETE("/tags/:id", v1.DeletedTag)

		GroupV1.GET("/articles", v1.GetArticles)
		GroupV1.GET("/articles/:id", v1.GetArticle)
		GroupV1.POST("/articles", v1.AddArticle)
		GroupV1.PUT("/articles/:id", v1.EditArticle)
		GroupV1.DELETE("/articles/:id", v1.DeleteArticle)
	}
}
