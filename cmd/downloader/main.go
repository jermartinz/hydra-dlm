package main

import (
	"log"

	"github.com/jermartinz/hydra-dlm/internal/downloader"
)

func main() {
	url := "https://golang.org/doc/gopher/frontpage.png"
	destination := "image.png"

	download, err := downloader.NewDownload(url, destination)
	if err != nil {
		log.Fatalf("Error creating download: %v", err)
	}

	err = download.Start()
	if err != nil {
		log.Fatalf("Error starting download: %v", err)
	}
	log.Println("Download completed successfully")

}
