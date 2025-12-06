package main

import (
	"bufio"
	"fmt"
	"log"
)

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
		b1, b1_idx, b2, b2_idx := 0, 0, 0, 0
		for i := 0; i < len(line); i++ {
			num := int(line[i] - '0')
			if b1 < num {
				if i == len(line)-1 {
					b2, b2_idx = b1, b1_idx
					b1, b1_idx = num, i
				} else {
					b2, b2_idx = 0, 0
					b1, b1_idx = num, i
				}
			} else if b2 < num {
				b2, b2_idx = num, i
			}
		}
		energy := 0
		if b1_idx < b2_idx {
			energy = ((b1 * 10) + b2)
		} else {
			energy = ((b2 * 10) + b1)
		}
		fmt.Printf("b1: %v, b2: %v, energy %v\n", b1, b2, energy)
		ret += energy
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(ret, "\n", err)
	}

	fmt.Println(ret)
}
