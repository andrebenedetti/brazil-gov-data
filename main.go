package main

import "gov-data/lib/discovery"

func main() {
	finder := discovery.GovBrCnpjFinder{}
	finder.FindFiles()
}
