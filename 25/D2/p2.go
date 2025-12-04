package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func check_range2(start string, end string, id_sum int, range_map map[int]*Range) int {
	i_start, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}
	var r Range
	r.start = i_start

	i_end, err := strconv.Atoi(end)
	if err != nil {
		log.Fatal(err)
	}
	r.end = i_end

	curr := i_start
	for curr < i_end {
		old_r, ok := range_map[curr]
		if ok && old_r.end < r.end {
			r.id_sum += old_r.id_sum
			r.invalids += old_r.invalids
			curr = old_r.end + 1
		}
		curr_str := strconv.Itoa(curr)
		if strings.Contains((curr_str + curr_str)[1:len(curr_str)], curr_str) {
			_, ok := range_map[i_start]
			if !ok {
				range_map[i_start] = &r
			}
			r.invalids++

			fmt.Println(" ", curr_str)
			id_sum += curr
			r.id_sum += curr
		}
		curr++
	}

	return id_sum
}

func Part2(filename string) {
	/*
	* read in file and split ranges into array of Ranges
	* log start value, start searching for invalid IDs
	* when invalid is found, log it and log range end
	* filter out all numbers with an odd length
	* */

	id_sum := 0
	range_map := make(map[int]*Range)
	for pair := range strings.SplitSeq(readline(filename), ",") {
		fmt.Println(pair)

		pair := strings.Split(pair, "-")
		start, end := pair[0], pair[1]
		len_s, len_e := len(start), len(end)
		if len_s == len_e && len_s&1 == 1 {
			continue
		}
		id_sum = check_range2(start, end, id_sum, range_map)
	}
	fmt.Println("\nID sum:", id_sum)
}
