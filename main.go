package main

import (
	"flag"
	"fmt"
	"log"
)

var banners = map[string]string{
	"standard":   "standard.txt",
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
}

func main() {
	checksum := flag.Bool("checksum", false, "Check integrity of specified file")
	outputFile := flag.String("output", "", "Output file name e.g output.txt")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		PrintError()
		return
	}

	input := args[0]
	var banner string

	if len(args) > 1 {
		banner = args[1]
	} else {
		banner = "standard"
	}

	filename, ok := banners[banner]
	if !ok {
		fmt.Println("Invalid banner specified.")
		return
	}

	if *checksum {
		err := ValidateFileChecksum(filename)
		if err != nil {
			log.Fatalf("Error checking integrity: %v", err)
		}
		fmt.Printf("Integrity check passed for file: %s\n", filename)
		return
	}

	err := ValidateFileChecksum(filename)
	if err != nil {
		log.Printf("Error downloading or validating file: %v", err)
		return
	}

	asciiArtGrid, err := ReadAscii(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	output, err := WriteArt(input, asciiArtGrid)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if *outputFile != "" {
		err = WriteAscii(*outputFile, output)
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
	} else {
		err = PrintArt(input, asciiArtGrid)
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}
}

func PrintError() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}
