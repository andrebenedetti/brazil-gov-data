package main

import "gov-data/feed"

func main() {
	feed.DownloadFiles(feed.GetAvailableFiles())
}
