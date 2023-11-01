package main

import (
	"fmt"
	"gov-data/lib/loaders"
	"gov-data/lib/storage"
	"log"
)

func main() {
	storage := storage.NewFileStorage(".", "cnaes.json")
	loader := loaders.NewCnaeLoader(storage)
	err := loader.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data loaded")

}
