package resources

import (
	"fmt"
	"log"

	"github.com/studio-b12/gowebdav"
)

func ListFiles(client *gowebdav.Client, path string) ([]string, error) {
	files, err := client.ReadDir(path)
	if err != nil {
		log.Printf("WebDAV Error: %v", err)
		return nil, err
	}
	fmt.Println(files)

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	return filenames, nil
}
