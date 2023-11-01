package main

import (
	"fmt"
	"gov-data/lib/loaders"
	"gov-data/lib/storage"
	"log"
)

func main() {
	storage := storage.NewFileStorage(".", "companies.json")
	loader := loaders.NewCompanyLoader(storage)
	err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data loaded")

}
