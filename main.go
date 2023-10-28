package main

import (
	"fmt"
	"gov-data/aggregator"
)

func main() {
	result := aggregator.ParseFile("./backup/cnaes.csv")
	fmt.Println(result[100])
	fmt.Println(result[101])

	aggregator.PrintAllRecords(result)
}
