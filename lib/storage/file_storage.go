package storage

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type FileStorage struct {
	directory string
	filename  string
}

func NewFileStorage(dir string, filename string) *FileStorage {
	return &FileStorage{
		directory: dir,
		filename:  filename,
	}
}

func (fs *FileStorage) Store(data interface{}) {
	if err := os.Mkdir(fs.directory, 0700); err != nil {
		if !strings.Contains(err.Error(), "exists") {
			log.Fatalln(err)
		}
	}

	file, err := os.Create(fs.directory + "/" + fs.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := json.Marshal(data)
	file.Write(output)
}
