package pipeline

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Download(fileUrl string) []byte {
	resp, err := http.Get(fileUrl)
	if err != nil {
		log.Fatalf("Failed to get url %s\n", fileUrl)
	}

	length := resp.ContentLength
	fmt.Println(length)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err.Error())
	}

	return body
}

func DownloadWithProgress(fileUrl string) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		log.Fatalf("Failed to get url %s\n", fileUrl)
	}

	length := resp.ContentLength
	fmt.Printf("Downloading file with length: %d\n", length)
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	bytesRead := 0
	lastLoggedValue := 0
	logPercentStep := 1
	fmt.Println("Progress: 0%")
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		bytesRead += n

		if bytesRead*100/int(length)-logPercentStep > lastLoggedValue {
			// not entirely accurate, but better UX
			lastLoggedValue += logPercentStep
			fmt.Printf("Progress: %d%% \n", lastLoggedValue)
		}
	}
}
