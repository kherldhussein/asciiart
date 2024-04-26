package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadASCIIMapFromFile(filename string) ([][]string, error) {
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
