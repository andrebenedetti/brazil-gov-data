package discovery

// Retrieve links from https://dadosabertos.rfb.gov.br/CNPJ/
// List of all brazilian companies

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

//	    Download speed from dadosabertos.rfb.gov.br is known to be
//		really really slow. If you intend to download it over and over again,
//		for example because you need to keep your database up-to-date,
type GovBrCnpjFinder struct {
}

func isZipFilename(filename string) bool {
	return strings.HasSuffix(filename, ".zip")
}

func isUpdatedAtField(field string) bool {
	r := regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}`)
	return r.Match([]byte(field))
}

func (finder *GovBrCnpjFinder) FindFiles() []RemoteFileMetadata {
	var cnpjDataUrl = "https://dadosabertos.rfb.gov.br/CNPJ/"
	c := colly.NewCollector()

	files := make([]RemoteFileMetadata, 0, 200)

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
			files = append(files, RemoteFileMetadata{cnpjDataUrl + filename, filename, updatedAt})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(cnpjDataUrl)

	return files
}
