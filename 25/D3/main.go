package main

import (
	"log"
	"os"
)

func Readfile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return file
}

func main() {
	filename := "puzzle.txt"
	filename = "small_puzzle.txt"
	Part1(filename)
}
