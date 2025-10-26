package webdav

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	uri := os.Getenv("WEBDAV_URL")

	auth, err := NewAuth()
	if err != nil {
		log.Fatalf("Auth error WebDAV: %v", err)
	}

	client := NewWebdavClient(uri, *auth)

	h := newHandler(client)

	router.GET("/connect", h.Connect)
	router.GET("/disconnect", h.Disconnect)
	router.GET("/list_files", h.ListFiles)
}
