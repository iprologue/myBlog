package router

import (
	"net/http"

	"github.com/iprologue/myBlog/pkg/setting"

	"github.com/gin-gonic/gin"
	_ "github.com/iprologue/myBlog/docs"
	"github.com/iprologue/myBlog/middleware/jwt"
	"github.com/iprologue/myBlog/middleware/logger"
	"github.com/iprologue/myBlog/pkg/upload"
	"github.com/iprologue/myBlog/router/api"
	v1 "github.com/iprologue/myBlog/router/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {

	engine := gin.New()
	gin.SetMode(setting.ServerSetting.RunMode)
	engine.Use(gin.Logger(), gin.Recovery(), logger.LoggerToFile())

	engine.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	engine.GET("/auth", api.GetAuth)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.POST("/upload", api.UploadImage)

	GroupV1 := engine.Group("/api/v1").Use(jwt.JWT())
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

	return engine
}
