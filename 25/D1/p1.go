package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rotate1(pos int, direction string, amount int) int {
	switch direction {
	case "L":
		return (pos - amount) % 100
	case "R":
		return (pos + amount) % 100
	default:
		panic("Rotate was given a direction other than L or R")
	}
}

func Part1(pos int, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	num_zero := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Read one line
		direction := string(rune(line[0]))
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		pos = rotate1(pos, direction, amount)
		if pos == 0 {
			num_zero++
		}
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Println(num_zero)
}
