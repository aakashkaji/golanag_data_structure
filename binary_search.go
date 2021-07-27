package main

import "fmt"

/* concept
left = 0
right = len(arr)-1
middle = (left +right) / 2

*/


func binarySearch(arr []int, target int) bool {

	left := 0
	right := len(arr)-1

	for left <= right {
		// calculate middle
		middle := (left+right)/2
		if arr[middle] == target {
			return true
		}else if arr[middle] > target {
			right = middle -1
		}else {
			left = middle + 1
		}

	}
	return false
}

func main() {

	a := []int {-2,-1,0,2,4,5,6,7,8,10,16,17,18,20}

	fmt.Println(binarySearch(a, 10))


}