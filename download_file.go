package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(file string) error {
	url := ""
	switch file {
	case "standard.txt":
		url = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"
	case "shadow.txt":
		url = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/shadow.txt"
	case "thinkertoy.txt":
		url = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/thinkertoy.txt"
	default:
		return fmt.Errorf("unsupported file name: %s", file)
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Write to file
	err = os.WriteFile(file, body, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Println("Downloaded", file, "from", url)
	return nil
}
