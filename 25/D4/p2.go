package main

import (
	"bufio"
	"fmt"
)

// func count_ats2(grid []string, row int, col int) int {
// 	DIRECTION := [][]int{
// 		{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1},
// 	}
// 	COL_LEN := len(grid[0])
// 	ROW_LEN := len(grid)
//
// 	ret := 0
//
// 	offset := 1
// 	for d := range DIRECTION {
// 		// for offset := 1; offset < COL_LEN; offset++ {
// 		curr_x := row + (offset * DIRECTION[d][0])
// 		curr_y := col + (offset * DIRECTION[d][1])
//
// 		if curr_x < 0 || ROW_LEN-1 < curr_x {
// 			continue
// 		} else if curr_y < 0 || COL_LEN-1 < curr_y {
// 			continue
// 		} else if grid[curr_x][curr_y] == '@' {
// 			ret++
// 			// }
// 		}
// 	}
// 	return ret
// }

type XY struct {
	x int
	y int
}

func Part2(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	var grid []string
	var rm_pos []XY

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

	ret, rnd := 0, -1
	for rnd != 0 {
		rnd = 0
		for i, row := range grid {
			for j := range row {
				if grid[i][j] == '.' {
					fmt.Printf(". ")
					continue
				}

				count := count_ats(grid, i, j)
				if count < 4 {
					rnd++
					fmt.Printf("X ")
					rm_pos = append(rm_pos, XY{i, j})
				} else {
					fmt.Printf("%v ", count)
				}
			}
			fmt.Println()
		}
		for _, pos := range rm_pos {
			line := grid[pos.x]
			grid[pos.x] = line[:pos.y] + "." + line[pos.y+1:]
		}
		rm_pos = rm_pos[:0]
		ret += rnd
	}
	fmt.Println(ret)
}
