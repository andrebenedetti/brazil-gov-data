package main

import (
	"gov-data/lib/loaders"
	"gov-data/lib/storage"
	"log"
)

func main() {
	loader := loaders.CompanyLoader{}
	data, err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewFileStorage(".", "companies.json")
	storage.Store(data)
}
