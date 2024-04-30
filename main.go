package main

import (
	"flag"
	"fmt"
	"log"
)

var formats = map[string]string{
	"s":  "standard.txt",
	"t":  "thinkertoy.txt",
	"sh": "shadow.txt",
}

func main() {
	formatPtr := flag.String("f", "s", "Output format: s (standard), t (thinkertoy), sh (shadow)")
	flag.Parse()

	filename, ok := formats[*formatPtr]
	if !ok {
		log.Fatalf("Invalid format specified: %s", *formatPtr)
	}

	asciMap, err := ReadASCIIMapFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please provide exactly one input string.")
		return
	}
	input := args[0]

	err = PrintArt(input, asciMap)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
