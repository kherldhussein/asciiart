package main

import (
	"fmt"
	"strings"
)

func WriteArt(str string, asciiArtGrid [][]string) (string, error) {
	var output strings.Builder

	switch str {
	case "":
		output.WriteString("")
	case "\\n":
		output.WriteString("\n")
	case "\\r", "\\f", "\\v", "\\t", "\\b", "\\a":
		return "", fmt.Errorf("error: unsupported escape sequence '%s'", str)
	default:
		s := strings.ReplaceAll(str, "\\n", "\n")
		s = strings.ReplaceAll(s, "\\r", "\r")
		s = strings.ReplaceAll(s, "\\f", "\f")
		s = strings.ReplaceAll(s, "\\v", "\v")
		s = strings.ReplaceAll(s, "\\t", "\t")
		s = strings.ReplaceAll(s, "\\b", "\b")
		s = strings.ReplaceAll(s, "\\a", "\a")
		words := strings.Split(s, "\n")
		num := 0
		for _, word := range words {
			if word == "" {
				num++
				if num < len(words) {
					output.WriteString("\n")
					continue
				}
			} else {
				err := printWordArt(&output, word, asciiArtGrid)
				if err != nil {
					return "", err
				}
			}
		}
	}
	return output.String(), nil
}

func printWordArt(output *strings.Builder, word string, asciiArtGrid [][]string) error {
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			if index < 0 || index >= len(asciiArtGrid) {
				return fmt.Errorf("unknown character: %q", char)
			} else {
				output.WriteString(asciiArtGrid[index][i])
			}
		}
		output.WriteString("\n")
	}
	return nil
}
