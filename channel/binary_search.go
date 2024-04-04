package main

import "fmt"

func binarySearch(arr []int, target int, low, high int) int {

	// based condition return from the function

	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if target == arr[mid] {
		return mid
	} else if arr[mid] < target {
		return binarySearch(arr, target, mid+1, high)
	} else {
		return binarySearch(arr, target, low, mid-1)
	}
}

func main() {

	//binary search
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 7

	fmt.Println("Index of target: ", binarySearch(num, target, 0, len(num)-1))

}
