package main

import (
	"fmt"
	"sort"
)

func main() {

	array_one := []int{7, 9, 50, 1, 2, 3, 4, 5, 6}

	array_two := []int{51, 5, 6, 8, 9} // sort this array
	sort.Ints(array_one)
	sort.Ints(array_two)

	i, j := 0, 0

	for i < len(array_one) && j < len(array_two) {

		if array_one[i] == array_two[j] {
			fmt.Println(array_one[i])
			i++
			j++
		} else if array_one[i] < array_two[j] {
			i++
		} else if array_one[i] > array_two[j] {
			j++

		}
	}

}
