package main

import (
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

	storage := storage.NewFileStorage(".", "cnaes.json")
	storage.Store(data)
}
