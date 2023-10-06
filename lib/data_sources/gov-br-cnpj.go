package data_sources

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var cnpjDataUrl = "https://dadosabertos.rfb.gov.br/CNPJ/"

func isZipFilename(filename string) bool {
	return strings.HasSuffix(filename, ".zip")
}

func isUpdatedAtField(field string) bool {
	r := regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}`)
	return r.Match([]byte(field))
}

func GetAvailableFiles() []FileSnapshot {
	c := colly.NewCollector()

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
			files = append(files, FileSnapshot{cnpjDataUrl + filename, filename, updatedAt})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(cnpjDataUrl)

	fmt.Println(files)
	return files
}
