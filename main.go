package main

import (
	"fmt"
	"gov-data/lib/data_loaders"
)

func main() {
	loader := data_loaders.CnaeLoader{}
	results, _ := loader.Load()
	fmt.Println(results)

}
