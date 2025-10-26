package resources

import (
	"github.com/studio-b12/gowebdav"
)

func Connect(client *gowebdav.Client) error {
	return client.Connect()
}

// func Disconnect(client *gowebdav.Client) error {
// 	return client.Disconnect()
// }
