package pipeline

import (
	"io"
	"log"
	"net/http"
)

func Download(fileUrl string) []byte {
	resp, err := http.Get(fileUrl)
	if err != nil {
		log.Fatalf("Failed to get url %s\n", fileUrl)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err.Error())
	}

	return body
}
