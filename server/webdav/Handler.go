package webdav

import (
	"net/http"
	"workbench/reiverrbridge/webdav/resources"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	client *WebdavClient
}

func newHandler(client *WebdavClient) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) Connect(c *gin.Context) {
	err := resources.Connect(h.client.Client)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "connected"})
}

func (h *Handler) Disconnect(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "disconnected"})
}

func (h *Handler) ListFiles(c *gin.Context) {
	reqPath := c.Query("path")
	if reqPath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter 'path' is missing"})
		return
	}
	if reqPath == "/" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Root folder is unauthorized to be listed"})
		return
	}

	filenames, err := resources.ListFiles(h.client.Client, reqPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files": filenames,
	})
}
