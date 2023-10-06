package data_sources

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type FileSnapshot struct {
	url       string
	filename  string
	updatedAt string
}

var filesDir = "gov-br"

func prepareDir() {
	if err := os.Mkdir(filesDir, 0755); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatalln(err)
		}
	}
}

func DownloadFiles(files []FileSnapshot) {
	prepareDir()

	for _, file := range files {
		fmt.Println(file)
		localFile, err := os.Create(filesDir + "/" + file.filename)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Get(file.url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		size, err := io.Copy(localFile, resp.Body)
		fmt.Printf("Got file size %d", size)

		defer localFile.Close()
	}
}
