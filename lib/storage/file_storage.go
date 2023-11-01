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
	file      *os.File
	format    string
}

func NewFileStorage(dir string, filename string) *FileStorage {

	if err := os.Mkdir(dir, 0700); err != nil {
		if !strings.Contains(err.Error(), "exists") {
			log.Fatalln(err)
		}
	}

	file, err := os.Create(dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	return &FileStorage{
		directory: dir,
		filename:  filename,
		file:      file,
		format:    "json",
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

func (fs *FileStorage) Write(data interface{}) error {
	// currently, we only support JSON format
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, writeErr := fs.file.Write(output)
	return writeErr
}

func (fs *FileStorage) Close() {
	if fs.file != nil {
		fs.file.Close()
	}
}
