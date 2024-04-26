package main

import (
	"fmt"
	"regexp"
	"strings"
)

func printWord(word string, asciMap [][]string) error {
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			if index < 0 || index >= len(asciMap) {
				return fmt.Errorf("unknown character: %c", char)
			} else {
				fmt.Print(asciMap[index][i])
			}
		}
		fmt.Println()
	}
	return nil
}

func PrintArt(str string, asciMap [][]string) error {
	escapes := map[string]string{
		"\\n": "\n",
		"\\r": "\r",
		"\\f": "\f",
		"\\v": "\v",
		"\\t": "\t",
		"\\b": "\b",
		"\\a": "\a",
	}

	escapesPattern := strings.Join(getKeys(escapes), "|")
	re := regexp.MustCompile(escapesPattern)

	s := re.ReplaceAllStringFunc(str, func(match string) string {
		return escapes[match]
	})

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
	return nil
}

func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, regexp.QuoteMeta(k))
	}
	return keys
}
