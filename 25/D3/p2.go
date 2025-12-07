package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
)

func Part2(filename string) {
	f := ReadFile(filename)
	defer f.Close()

	const window_size = 12
	ret := 0
	line_num := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line_num++
		fmt.Println("Line", line_num)
		line := scanner.Text() // Read one line

		min_heap := &IntMinHeap{}
		heap.Init(min_heap)
		var arr [window_size]int

		fmt.Println(line)
		for i := range window_size {
			num := int(line[i] - '0')
			heap.Push(min_heap, num)
			arr[i] = num
		}

		for i := window_size; i < len(line); i++ {
			num := int(line[i] - '0')
			min_num := min_heap.Peek()
			if num > min_num {
				min_heap.Put(0, num)

				idx := 0
				for j := range arr {
					if arr[j] == min_num {
						idx = j
						break
					}
				}
				for j := idx; j < len(arr)-1; j++ {
					arr[j] = arr[j+1]
				}
				arr[len(arr)-1] = num
			}
		}
		fmt.Printf("Min: %v\nHeap: %v\nArr:   %v\n", min_heap.Peek(), min_heap, arr)
		sum := 0
		for _, num := range arr {
			sum = sum*10 + num
		}
		fmt.Printf("Sum: %v\n\n", sum)
		ret += sum
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(ret, "\n", err)
	}

	fmt.Println(ret)
}

/* INFO:
*  Line 1
*  987654321111 c
*  987654321111 c
*  Line 2
*  811111111119 c
*  811111111119 c
*  Line 3
*  343423423478 x
*  4 34234234278 c
*  Line 4
*  818181911112 x
*  888911112111 c
* */
