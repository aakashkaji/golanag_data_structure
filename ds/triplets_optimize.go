package main

import (
	"fmt"
	"sort"
)

func main() {

	var triplets [][]int
	arr := []int{12, 3, 6, 1, 6, 9, 6}
	sum := 24

	sort.Ints(arr)
	
	for i := 0; i < len(arr)-2; i++ {

		left, right := i+1, len(arr)-1

		for left < right {

			target := arr[i] + arr[left] + arr[right]

			if sum == target {
				triplets = append(triplets, []int{arr[i], arr[left], arr[right]})

				left++
				right--
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}

			} else if sum < target {
				right--
			} else {
				left++
			}
		}
	}

	fmt.Println(triplets)

}
