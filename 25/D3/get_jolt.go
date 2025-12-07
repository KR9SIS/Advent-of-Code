package main

import (
	"fmt"
	"math"
)

func remove_idx(i int, rm_idx *[12]int, sum int) int {
	s1 := sum
	s1 = s1 - rm_idx[i]
	// fmt.Println("sum:", sum, "\nrm i", s1)

	for j := i - 1; j >= 0; j-- {
		s1 = s1 - rm_idx[j] + (rm_idx[j] / 10)
		// fmt.Println("rm j", s1)
	}
	return (s1 * 10)
}

func rm_arr_idx(i int, rm_idx *[12]int, new_num int) {
	for j := i; j < len(rm_idx)-1; j++ {
		rm_idx[j] = rm_idx[j+1] * 10
	}
	rm_idx[len(rm_idx)-1] = new_num
}

func GetJolt(line string) int {
	const win_size = 12

	// var arr [win_size]int
	var rm_idx [win_size]int

	sum := 0
	for i := range win_size {
		num := int(line[i] - '0')
		// arr[i] = num
		sum = sum*10 + num
		rm_idx[i] = (num * (int(math.Pow10(win_size - i - 1))))
	}

	tmp_sum := sum

	for c := win_size; c < len(line); c++ {
		new_num := int(line[c] - '0')

		for i := range win_size {

			removed_i := remove_idx(i, &rm_idx, tmp_sum)
			new_sum := removed_i + new_num

			fmt.Println(" Tmp:", tmp_sum, "\nCalc:", new_sum)
			fmt.Println()

			if new_sum > tmp_sum {
				tmp_sum = new_sum
				rm_arr_idx(i, &rm_idx, new_num)
				// rm_idx[i] = (new_num * (int(math.Pow10(win_size - i - 1))))
			}
		}
		fmt.Println("Finish ", c)
	}

	fmt.Println("\nFinal Sum:", tmp_sum)
	return sum
}
