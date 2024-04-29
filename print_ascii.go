package main

import (
	"fmt"
	"strings"
)

func printWord(word string, asciMap [][]string) error {
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			if index < 0 || index >= len(asciMap) {
				return fmt.Errorf("unknown character: %q", char)
			} else {
				fmt.Print(asciMap[index][i])
			}
		}
		fmt.Println()
	}
	return nil
}

func PrintArt(str string, asciMap [][]string) error {
	switch str {
	case "":
		fmt.Print()
	case "\\n":
		fmt.Println()
	case "\\r", "\\f", "\\v", "\\t", "\\b", "\\a":
		return fmt.Errorf("error: unsupported escape sequence '%s'", str)
	default:
		s := strings.ReplaceAll(str, "\\n", "\n")
		s = strings.ReplaceAll(s, "\\r", "\r")
		s = strings.ReplaceAll(s, "\\f", "\f")
		s = strings.ReplaceAll(s, "\\v", "\v")
		s = strings.ReplaceAll(s, "\\t", "\t")
		s = strings.ReplaceAll(s, "\\b", "\b")
		s = strings.ReplaceAll(s, "\\a", "\a")
		words := strings.Split(s, "\n")
		for _, word := range words {
			if word == "" {
				fmt.Println()
				continue
			}
			err := printWord(word, asciMap)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
