package webdav

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	uri := os.Getenv("WEBDAV_URL")

	client := NewWebdavClient(uri)

	h := newHandler(client)

	router.GET("/connect", h.Connect)
	router.GET("/disconnect", h.Disconnect)
	router.GET("/list_files", h.ListFiles)
	// router.GET("/stream", h.SocketStream)
	router.GET("/stream_media", h.SocketStream)
	router.GET("/get_media_path", h.GetMediaPath)
	router.GET("/get_media_token", h.GetMediaToken)
}
