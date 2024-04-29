package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	asciMap, err := ReadASCIIMapFromFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading ASCII map:", err)
		return
	}
	if len(args) < 1 {
		fmt.Println("Please provide an input string.")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments passed")
		return
	}
	err = PrintArt(args[0], asciMap)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}
}
