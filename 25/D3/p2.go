package main

import (
	"bufio"
	"fmt"
	"log"
)

func get_max_digit(line string) int {
	const win_size = 12
	const line_len = 100
	var m [line_len + 1][win_size + 1]int

	for i := range line_len {
		number := int(line[i] - '0') // Get character at index i as int

		max_single_digit := m[i][1]
		m[i+1][1] = max(max_single_digit, number)

		for j := 2; j <= win_size; j++ {
			current_max_combined := m[i][j-1]*10 + number
			previous_max := m[i][j]
			m[i+1][j] = max(current_max_combined, previous_max)
		}
	}

	return m[line_len][win_size]
}

func Part2(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	ret := 0
	line_num := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line_num++
		fmt.Println("Line", line_num)
		line := scanner.Text() // Read one line

		line_ret := get_max_digit(line)
		fmt.Println(line_ret)
		ret += line_ret
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(ret, "\n", err)
	}

	fmt.Println("\nResult:", ret)
}
