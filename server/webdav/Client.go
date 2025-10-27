package webdav

import (
	"os"

	"github.com/studio-b12/gowebdav"
)

type WebdavClient struct {
	Client *gowebdav.Client
}

func NewWebdavClient(uri string) *WebdavClient {
	username := os.Getenv("WEBDAV_USERNAME")
	password := os.Getenv("WEBDAV_PASSWORD")
	c := gowebdav.NewClient(uri, username, password)

	return &WebdavClient{
		Client: c,
	}
}
