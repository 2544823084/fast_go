package router

import (
	"net/http"

	"gobox/framwork/gin/internal"
	"gobox/framwork/gin/middleware"

	"github.com/gin-gonic/gin"
)

// Setup 创建并配置 Gin 引擎（中间件 + 路由）。
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORS(), middleware.Logger(), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/file/upload", internal.UploadFile)
	r.POST("/file/download", internal.DownloadFile)
	r.POST("/file/delete", internal.DeleteFile)
	r.POST("/file/list", internal.GetFileList)

	return r
}
