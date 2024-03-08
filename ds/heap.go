package main

import "fmt"

func minHeapify(arr []int, j, i int) {
	n := 3
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}

	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		minHeapify(arr, n, smallest)

	}

}

func maxHeapify(arr []int, n, i int) {

	gratest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[gratest] {
		gratest = left
	}
	if right < n && arr[right] > arr[gratest] {
		gratest = right
	}

	if gratest != i {
		arr[i], arr[gratest] = arr[gratest], arr[i]
		maxHeapify(arr, n, gratest)
	}

	fmt.Println(arr)
}

func heapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(arr, n, i)
		fmt.Println("call hippp")
	}

	// Heap sort
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		minHeapify(arr, i, 0)
	}
}

func main() {
	arr := []int{0, 11, 12, 13, 15, 6, 7, 60}
	fmt.Println("Original array:", arr)

	heapSort(arr)

	fmt.Println("Sorted array:", arr)
}
