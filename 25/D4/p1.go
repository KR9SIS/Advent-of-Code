package main

import (
	"bufio"
	"fmt"
)

func count_ats(grid []string, row int, col int) int {
	DIRECTION := [][]int{
		{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1},
	}
	COL_LEN := len(grid[0])
	ROW_LEN := len(grid)

	ret := 0

	offset := 1
	for d := range DIRECTION {
		// for offset := 1; offset < COL_LEN; offset++ {
		curr_x := row + (offset * DIRECTION[d][0])
		curr_y := col + (offset * DIRECTION[d][1])

		if curr_x < 0 || ROW_LEN-1 < curr_x {
			continue
		} else if curr_y < 0 || COL_LEN-1 < curr_y {
			continue
		} else if grid[curr_x][curr_y] == '@' {
			ret++
			// }
		}
	}
	return ret
}

func Part1(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	var grid []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	for _, row := range grid {
		for i := range row {
			fmt.Printf("%v ", string(row[i]))
		}
		fmt.Println()
	}
	fmt.Println()

	ret := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '.' {
				continue
			}

			count := count_ats(grid, i, j)
			if count < 4 {
				ret++
				fmt.Printf("X ")
			} else {
				fmt.Printf("%v ", count)
			}
		}
		fmt.Println()
	}

	fmt.Println(ret)
}
