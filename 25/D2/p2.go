package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func check_range2(start string, end string, id_sum int) int {
	i_start, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}

	i_end, err := strconv.Atoi(end)
	if err != nil {
		log.Fatal(err)
	}

	curr := max(i_start, 11)
	for curr <= i_end {
		curr_str := strconv.Itoa(curr)
		str_cat := (curr_str + curr_str)[1:]
		if strings.Index(str_cat, curr_str) < len(curr_str)-1 {
			fmt.Println(" ", curr_str)
			id_sum += curr
		}
		curr++
	}

	return id_sum
}

func Part2(filename string) {
	id_sum := 0
	for pair := range strings.SplitSeq(readline(filename), ",") {
		fmt.Println(pair)

		pair := strings.Split(pair, "-")
		start, end := pair[0], pair[1]
		id_sum = check_range2(start, end, id_sum)
	}
	fmt.Println("\nID sum:", id_sum)
}
