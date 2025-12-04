package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rotate2(pos int, direction string, amount int) (p int, z int) {
	fmt.Printf("p: %v, d: %v, a: %v\n", pos, direction, amount)
	zeroes := amount / 100
	if zeroes != 0 {
		amount = amount - (100 * zeroes)
	}
	init_pos := pos

	switch direction {
	case "L":
		pos = pos - amount
		if pos < 0 && init_pos != 0 {
			pos = pos + 100
			zeroes++
		}
	case "R":
		pos = pos + amount
		if pos > 99 {
			zeroes++
		}
	default:
		panic("Rotate was given a direction other than L or R")
	}

	pos = pos % 100
	fmt.Printf("p: %v z: %v\n\n", pos, zeroes)
	return pos, zeroes
}

func Part2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	num_zero := 0

	pos := 50
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Read one line
		direction := string(rune(line[0]))
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		zeroes := 0
		pos, zeroes = rotate2(pos, direction, amount)
		pos = pos + 0

		if zeroes != 0 {
			num_zero += zeroes
		}
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Println(num_zero)
}
