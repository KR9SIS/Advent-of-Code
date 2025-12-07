package main

import (
	"bufio"
	"fmt"
	"log"
)

func Part2(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	const window_size = 12
	ret := 0
	line_num := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line_num++
		fmt.Println("Line", line_num)
		line := scanner.Text() // Read one line

		ret += GetJolt(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(ret, "\n", err)
	}

	fmt.Println(ret)
}

/* INFO:
*  Line 1
*  987654321111 c
*  987654321111 c
*  Line 2
*  811111111119 c
*  811111111119 c
*  Line 3
*  343423423478 x
*  4 34234234278 c
*  Line 4
*  818181911112 x
*  888911112111 c
* */
