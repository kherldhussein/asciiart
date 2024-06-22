package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var files = map[string]bool{
	"shadow.txt":     true,
	"standard.txt":   true,
	"thinkertoy.txt": true,
}

func validateFileName(file string) bool {
	_, ok := files[file]
	return ok
}

func ReadASCIIMapFromFile(filename string) ([][]string, error) {
	if !validateFileName(filename) {
		return nil, fmt.Errorf("unsupported file name: %s", filename)
	}

	if !strings.HasSuffix(filename, ".txt") {
		return nil, fmt.Errorf("unsupported file format: %s", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()
	var (
		asciMap  [][]string
		asciLine []string
	)

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		lines := scanner.Text()
		asciLine = append(asciLine, lines)
		count++
		if count == 9 {
			asciMap = append(asciMap, asciLine)
			count = 0
			asciLine = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}
	return asciMap, nil
}
