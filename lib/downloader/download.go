package downloader

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func Download(url string, writer io.Writer, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get url %s\n", url)
	}

	defer resp.Body.Close()
	io.Copy(writer, resp.Body)
	wg.Done()
}
