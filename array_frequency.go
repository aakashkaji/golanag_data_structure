package main

import "fmt"

var hashmap map[int]int

func main() {

	a := []int{1, 2, 3, 4, 1, 2, 1, 3, 4, 7, 8, 9, 0}

	hashmap = make(map[int]int)

	for i := 0; i < len(a); i++ {

		if _, ok := hashmap[a[i]]; ok {
			hashmap[a[i]]++
		} else {
			hashmap[a[i]] = 1
		}
	}

	// hashmap sort by values

	fmt.Println(hashmap, a)
}
