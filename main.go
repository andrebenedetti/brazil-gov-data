package main

import "gov-data/lib/data_sources"

func main() {
	data_sources.DownloadFiles(data_sources.GetAvailableFiles())
}
