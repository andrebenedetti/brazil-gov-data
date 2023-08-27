package feed

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var cnpjDataUrl = "https://dadosabertos.rfb.gov.br/CNPJ/"

func prepareDir() {
	if err := os.Mkdir("data", 0755); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatalln(err)
		}
	}
}

type FileSnapshot struct {
	filename  string
	updatedAt string
}

func isZipFilename(filename string) bool {
	return strings.HasSuffix(filename, ".zip")
}

func isUpdatedAtField(field string) bool {
	r := regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}`)
	return r.Match([]byte(field))
}

func GetAvailableFiles() []string {
	c := colly.NewCollector()

	// Find and visit all links
	// c.OnHTML("td a[href]", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Attr("href"))
	// })

	files := make([]FileSnapshot, 0, 200)

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		values := e.ChildTexts("td")
		var filename string
		var updatedAt string
		for _, value := range values {
			fmt.Println(value)
			if isZipFilename(value) {
				filename = value
			} else if isUpdatedAtField(value) {
				updatedAt = value
			}
		}

		if filename != "" && updatedAt != "" {
			files = append(files, FileSnapshot{filename, updatedAt})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(cnpjDataUrl)

	fmt.Println(files)
	return []string{}
}

// resp, err := http.Get(url)
// if err != nil {
// 	log.Fatal(err)
// }
// defer resp.Body.Close()

// b, err := io.ReadAll(resp.Body)

// if err != nil {
// 	log.Fatalln(err)
// }

// for _, match := range matches {
// 	filename := strings.Split(match[0], "\"")[1]

// 	file, err := os.Create("./data/" + filename)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Put content on file
// 	resp, err := http.Get(url + filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	size, err := io.Copy(file, resp.Body)
// 	fmt.Printf("Got file size %d", size)

// 	defer file.Close()
// }
