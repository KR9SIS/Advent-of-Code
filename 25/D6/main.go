package main

import (
	"log"
	"os"
)

func ReadFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	filename := "puzzle.txt"
	filename = "small_puzzle.txt"
	Part2(filename)
}
