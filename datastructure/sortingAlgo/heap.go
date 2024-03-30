package main

import "fmt"

func buildHeap(nums []int) {

	n := len(nums)
	for i := (n / 2) - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	fmt.Println(nums)

}

func heapify(arr []int, n, i int) {

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
		heapify(arr, n, gratest)
	}
}

func main() {
	arr := []int{0, 11, 12, 13, 15, 6, 7, 60}

	buildHeap(arr)

	fmt.Println("Sorted array:", arr)
}
