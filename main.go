package main

import (
	"fmt"
	"gov-data/lib/loaders"
	"gov-data/lib/storage"
)

func main() {
	store := storage.NewPostgresStore()
	loader := loaders.NewCnaeLoader(store)
	loader.Load()
	result := store.GetCnaes()
	fmt.Println(result.Next())
	fmt.Println(result.Values())
}
