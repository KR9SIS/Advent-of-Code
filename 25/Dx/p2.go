package main

import (
	"log"
	"os"
)

func Part2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
