package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	fmt.Println("--------", arr)
	left := 0
	right := len(arr) - 1
	pivot := arr[(len(arr)-1)/2]

	for left <= right {

		for pivot > arr[left] {
			left++
		}
		for pivot < arr[right] {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
			fmt.Println("-----swap")
		}
		fmt.Println("***", arr)

	}

	fmt.Println(arr[:left], arr[left:])
	quickSort(arr[:left])
	quickSort(arr[left:])

}

func main() {

	arr := []int{4, 8, 24, 7, 5, 0, 36, 32, 30, 3}
	quickSort(arr)
	fmt.Println(arr)

}
