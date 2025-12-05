package main

import (
	"bufio"
	"log"
	"os"
)

func readline(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	}
	log.Fatal("scanner.Scan returned false")
	return ""
}

func count_digits(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}

func main() {
	filename := "puzzle.txt"
	// filename = "small_puzzle.txt"
	Part2(filename)
}
