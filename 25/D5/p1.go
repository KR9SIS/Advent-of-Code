package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func Part1(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	var ranges []Range
	ret := 0

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
		// fmt.Println(line)
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.start == b.start {
			// put larger range first
			return b.end - a.end
		}
		return a.start - b.start
	})
	fmt.Println(len(ranges))

	for scanner.Scan() {
		line := scanner.Text()

		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(id)
		idx, _ := slices.BinarySearchFunc(ranges, Range{id, id}, func(a, b Range) int {
			return a.start - b.start
		})
		if idx != 0 {
			// fmt.Println(ranges[idx : idx+3])
			idx--
		}
		for i := idx; i < idx+3; i++ {
			r := ranges[idx]
			if r.start <= id && id <= r.end {
				fmt.Println("ID:", id, "\nin r:", ranges[idx])
				ret++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fresh IDs:", ret)
}
