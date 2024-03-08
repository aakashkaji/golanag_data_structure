package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{1, 3, 4, 6, 5, 8, 9}
	hash_map := make(map[int]bool)
	var result [][]int

	// find the pair of number whose sum is 10
	//d = 10-num

	for _, val := range num {

		diff := 10 - val

		if _, ok := hash_map[diff]; ok {
			result = append(result, []int{val, diff})

		}
		hash_map[val] = true
	}

	fmt.Println(hash_map)
	fmt.Println(result)
	sort.Ints(num)
	fmt.Println(num)

}
