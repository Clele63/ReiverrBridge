package resources

import (
	"io"
	"log"

	"github.com/studio-b12/gowebdav"
)

func streamFromFile(client *gowebdav.Client, path string) (io.ReadCloser, error) {
	stream, err := client.ReadStream(path)
	if err != nil {
		log.Printf("WebDAV Error: %v", err)
		return nil, err
	}

	return stream, nil
}
