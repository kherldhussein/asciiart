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
	checkIntegrity := flag.Bool("checksum", false, "Check integrity of specified file")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Please provide at least one argument.")
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

	if *checkIntegrity {
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

	asciiMap, err := ReadASCIIMapFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	err = PrintArt(input, asciiMap)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
