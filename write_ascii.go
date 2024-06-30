package main

import (
	"io"
	"os"
	"strings"
)

func WriteAscii(filename, content string) error {
	if !strings.HasSuffix(filename, ".txt") {
		PrintError()
		return nil
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	if err != nil {
		return err
	}

	return nil
}
