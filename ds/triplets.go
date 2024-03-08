package main

import "fmt"

// Input : array = {12, 3, 6, 1, 6, 9} sum = 24
// Output : [[3, 9, 12], [6, 6, 12]]

func main() {

	arr := []int{12, 3, 6, 1, 6, 9}

	sum := 24

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {

			for k := j + 1; k < len(arr); k++ {

				if (arr[i] + arr[j] + arr[k]) == sum {
					fmt.Println([]int{arr[i], arr[j], arr[k]})
				}
			}
		}

	}

}