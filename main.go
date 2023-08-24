package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	url := "https://dadosabertos.rfb.gov.br/CNPJ/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	r := regexp.MustCompile(`href="(?P<link>.*.zip)"`)
	matches := r.FindAllStringSubmatch(string(b), -1)

	if err := os.Mkdir("data", 0755); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatalln(err)
		}
	}

	for _, match := range matches {
		filename := strings.Split(match[0], "\"")[1]

		file, err := os.Create("./data/" + filename)
		if err != nil {
			log.Fatalln(err)
		}

		// Put content on file
		resp, err := http.Get(url + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		size, err := io.Copy(file, resp.Body)
		fmt.Printf("Got file size %d", size)

		defer file.Close()
	}

}
