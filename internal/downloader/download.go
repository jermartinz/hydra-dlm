package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Download struct {
	URL             string
	Destination     string
	TotalBytes      int64
	DownloadedBytes int64
	Status          string // (pending, in_progress, pause, completed,  error)
	Error           error
}

func NewDownload(url, destination string) (*Download, error) {
	if url == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}
	if destination == "" {
		return nil, fmt.Errorf("Destination cannot be empty")
	}
	return &Download{
		URL:         url,
		Destination: destination,
		Status:      "pending",
	}, nil
}

func (d *Download) Start() error {

	outFile, err := os.Create(d.Destination)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	resp, err := http.Get(d.URL)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}
	return nil
}
