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
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})

	for _, r := range ranges {
		fmt.Println(r)
	}
	for scanner.Scan() {
		line := scanner.Text()

		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		for i := range ranges {
			r := ranges[i]
			if r.start <= id && id <= r.end {
				ret++
				break
			}
		}
	}
	fmt.Println("Fresh IDs:", ret)
}
