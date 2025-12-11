package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part2(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	setup := true
	problems_m := [][]int{}
	problem_ops := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		line_arr := strings.Fields(line)
		for i := range line_arr {
			str_num := line_arr[i]
			num, err := strconv.Atoi(str_num)
			if err != nil {
				problem_ops = append(problem_ops, line_arr[i])
			} else {
				if setup {
					problems_m = append(problems_m, []int{})
				}
				problems_m[i] = append(problems_m[i], num)
			}
		}
		setup = false
	}
	ret := 0
	for i, val := range problems_m {
		sum := 0
		switch problem_ops[i] {
		case "+":
			for _, num := range val {
				sum += num
			}
			ret += sum
		case "*":
			sum = 1
			for _, num := range val {
				sum *= num
			}
			ret += sum
		}

		fmt.Println(val, ":", problem_ops[i], "==", sum)
	}
	fmt.Println("\nGrand total:", ret)
}
