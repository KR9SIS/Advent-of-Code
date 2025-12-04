package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func check_range1(start string, end string, id_sum int, range_map map[int]*Range) int {
	var r Range
	i_start, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}
	r.start = i_start

	i_end, err := strconv.Atoi(end)
	if err != nil {
		log.Fatal(err)
	}
	r.end = i_end

	curr := 1
	s_len := len(start)
	if s_len > 1 {
		curr, err = strconv.Atoi(start[:s_len/2])
		if err != nil {
			log.Fatal(err)
		}
	}

	for curr < i_end {
		d_count := count_digits(curr)
		doubler := math.Pow10(d_count) + 1
		invalid_id := curr * int(doubler)

		if i_start <= invalid_id && invalid_id <= i_end {
			_, ok := range_map[i_start]
			if !ok {
				range_map[i_start] = &r
			}
			fmt.Println(invalid_id)
			r.invalids++
			id_sum += invalid_id
			r.id_sum += invalid_id
		}
		curr++

	}
	return id_sum
}

func Partq(filename string) {
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
		id_sum = check_range1(start, end, id_sum, range_map)
	}
	fmt.Println(id_sum)
}
