package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

// implement the heap interface
func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

}

func findTheLargestElement(arr []int, k int) int {
	h := &MinHeap{}

	heap.Init(h)

	for _, num := range arr {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}

	}

	return (*h)[0]

}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}

	k := 3
	fmt.Println(findTheLargestElement(nums, k))

}
