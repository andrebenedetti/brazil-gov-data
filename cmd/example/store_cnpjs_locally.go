package example

import (
	"fmt"
	"gov-data/lib/discovery"
	"gov-data/lib/downloader"
	"gov-data/lib/storage"
	"log"
	"sync"
)

func Run() {
	fileFinder := discovery.GovBrCnpjFinder{}
	sources := fileFinder.FindFiles()

	fileStorage := storage.FileStorage{Directory: "./data"}
	fmt.Printf("Found %d sources...", len(sources))

	var wg sync.WaitGroup
	for _, s := range sources {
		wg.Add(1)
		go func(src discovery.RemoteFileMetadata) {
			fmt.Println("Downloading ", src.Filename)
			file, err := fileStorage.OpenFile("./" + src.Filename)
			if err != nil {
				log.Fatalf("Error %s", err.Error())
			}
			defer file.Close()
			data := downloader.Download(src.Url)
			file.Write(data)
			wg.Done()
		}(s)
	}
	wg.Wait()
}
