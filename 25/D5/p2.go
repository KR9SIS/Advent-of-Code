package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func check_in_range(start, mid, end int) bool {
	if start < mid && mid < end {
		return true
	}
	return false
}

func Part2(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	var ranges []Range

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		pair := strings.Split(line, "-")
		s, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatal(err)
		}
		e, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, Range{s, e})
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})

	for i := 0; i < len(ranges)-1; i++ {
		modified := false
		r1 := ranges[i]
		r2 := ranges[i+1]

		if check_in_range(r1.start, r2.start, r1.end) || r1.end == r2.start {
			r2.start = r1.start
			modified = true
		}
		if check_in_range(r1.start, r2.end, r1.end) {
			r2.end = r1.end
			modified = true
		}
		if r1.start == r2.start {
			if r1.end > r2.end {
				r2 = r1
			}
			modified = true
		}

		if modified {
			ranges[i+1] = r2
			ranges[i] = Range{}
			modified = false
		}
	}

	num := 0
	ret := 0
	for _, r := range ranges {
		if r.start == 0 && r.end == 0 {
			continue
		}
		fmt.Println(r, ":", r.end-r.start+1)
		ret += r.end - r.start + 1
		num++
	}
	fmt.Println("\nFresh IDs:", ret, "\nNum ranges:", num)
}
