package pipeline

import (
	"io"
	"log"
	"net/http"
)

func Download(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get url %s\n", url)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err.Error())
	}

	return body
}
