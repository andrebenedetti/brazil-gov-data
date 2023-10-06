package storage

import (
	"log"
	"os"
	"strings"
)

type FileSnapshot struct {
	url       string
	filename  string
	updatedAt string
}

var filesDir = "gov-br"

type FileStorage struct {
	Directory string
}

func (fs *FileStorage) prepareDir() {
	if err := os.Mkdir(fs.Directory, 0700); err != nil {
		if !strings.Contains(err.Error(), "exists") {
			log.Fatalln(err)
		}
	}
}

func (fs *FileStorage) OpenFile(filename string) (*os.File, error) {
	fs.prepareDir()
	return os.Create(fs.Directory + "/" + filename)

}
