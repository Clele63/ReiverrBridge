package webdav

import (
	"os"

	"github.com/studio-b12/gowebdav"
)

type WebdavAuth struct {
	Authorizer gowebdav.Authorizer
}

type WebdavClient struct {
	Client *gowebdav.Client
	Auth   *WebdavAuth
}

func NewAuth() (*WebdavAuth, error) {
	auth := WebdavAuth{}

	username := os.Getenv("WEBDAV_USERNAME")
	password := os.Getenv("WEBDAV_PASSWORD")

	auth.Authorizer = gowebdav.NewAutoAuth(username, password)
	return &auth, nil
}

func NewWebdavClient(uri string, auth WebdavAuth) *WebdavClient {
	c := gowebdav.NewAuthClient(uri, auth.Authorizer)

	return &WebdavClient{
		Client: c,
		Auth:   &auth,
	}
}
