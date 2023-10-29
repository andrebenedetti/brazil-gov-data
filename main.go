package main

import (
	"encoding/json"
	"gov-data/lib/loaders"
	"gov-data/lib/storage"
	"log"
)

func main() {
	loader := loaders.CnaeLoader{}
	data, err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.FileStorage{}
	storage.Directory = "."
	file, err := storage.OpenFile("cnae.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	file.Write(output)
}
