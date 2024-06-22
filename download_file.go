package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var ExpectedSize = map[string]int64{
	"standard.txt":   6623,
	"shadow.txt":     7463,
	"thinkertoy.txt": 5558,
}

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
	// Check and handle corrupted file
	err := validateFileSize(file)
	if err != nil {
		return fmt.Errorf("error checking and handling corrupted file: %w", err)
	}

	// Proceed with download if file doesn't exist or is corrupted
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

func validateFileSize(file string) error {
	if _, err := os.Stat(file); err == nil {
		stat, err := os.Stat(file)
		if err != nil {
			return fmt.Errorf("failed to get file stat: %w", err)
		}
		if size, ok := ExpectedSize[file]; ok {
			if stat.Size() == size {
				fmt.Println("File", file, "already exists and is valid.")
				return nil
			} else {
				fmt.Println("File", file, "exists but is corrupted (size mismatch), deleting...")
				if err := os.Remove(file); err != nil {
					return fmt.Errorf("failed to delete corrupted file: %w", err)
				}
			}
		}
	}
	return nil
}
