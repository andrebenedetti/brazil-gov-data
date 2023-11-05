package main

import (
	"fmt"
	"gov-data/lib/storage"
)

func main() {
	store := storage.NewPostgresStore()
	result := store.GetCnaes()
	fmt.Println(result.Next())
	fmt.Println(result.Values())
}
