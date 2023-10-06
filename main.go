package main

import (
	"gov-data/lib/data_sources"
	"gov-data/lib/storage"
	"log"
	"sync"
)

func main() {
	fs := storage.FileStorage{Directory: "./data"}
	file, err := fs.OpenFile("sample.zip")
	if err != nil {
		log.Fatal("Error opening file")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	data_sources.Download("https://uol.com.br", file, &wg)
	wg.Wait()

}
