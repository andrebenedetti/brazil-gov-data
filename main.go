package main

import (
	"fmt"
	"gov-data/lib/loaders"
)

func main() {
	loader := loaders.CnaeLoader{}
	results, _ := loader.Load()
	fmt.Println(results)

}
