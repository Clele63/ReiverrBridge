package webdav

import (
	"net/http"
	"workbench/reiverrbridge/webdav/resources"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	client *WebdavClient
	cache  *resources.CacheToken
}

func newHandler(client *WebdavClient) *Handler {
	return &Handler{
		client: client,
		cache: &resources.CacheToken{
			MediaCache: make(map[string]string),
		},
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

	filenames, err := resources.ListFiles(h.client.Client, reqPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files": filenames,
	})
}

func (h *Handler) GetMediaPath(c *gin.Context) {
	searchText := c.Query("search")

	if searchText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le paramètre 'search' est manquant"})
		return
	}

	mediaPath, err := resources.GetMediaPath(h.client.Client, searchText)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"path": mediaPath,
	})
}

func (h *Handler) GetMediaToken(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Le paramètre 'path' est manquant"})
		return
	}
	// if path == "/" {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "L'accès à la racine n'est pas autorisé"})
	// 	return
	// }

	token := resources.GenerateToken(h.cache, path)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// func (h *Handler) StreamFile(c *gin.Context) {
// 	reqPath := c.Query("path")
// 	if reqPath == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter 'path' is missing"})
// 		return
// 	}

// 	stream, err := resources.SocketStream(c, h.client.Client)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.Header("Content-Type", "application/octet-stream")

// 	c.Stream(func(w io.Writer) bool {
// 		_, err := io.Copy(w, stream)

//			if err != nil {
//				log.Printf("Error while stream : %v", err)
//				return false
//			}
//			return true
//		})
//	}

func (h *Handler) SocketStream(c *gin.Context) {
	resources.SocketStream(c, h.client.Client, h.cache)
}
