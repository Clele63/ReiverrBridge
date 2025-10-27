package routes

import (
	"workbench/reiverrbridge/media"
	"workbench/reiverrbridge/webdav"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")

	webdavGroup := v1.Group("/webdav")
	{
		webdav.RegisterRoutes(webdavGroup)
	}
	mediaGroup := v1.Group("/media")
	{
		media.RegisterRoutes(mediaGroup)
	}
}
