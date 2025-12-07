package main

import (
	"container/heap"
)

// An IntMinHeap is a min-heap of ints.
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntMinHeap) Peek() int {
	return h[0]
}

func (h IntMinHeap) Index(i int) int {
	return h[i]
}

func (h IntMinHeap) Put(i, data int) {
	h[i] = data
	heap.Fix(&h, i)
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
/*
func test() {
	h := &IntMinHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
*/
