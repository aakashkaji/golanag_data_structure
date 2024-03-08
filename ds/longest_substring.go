package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func LongestSubstring() {

	s := "abcabcbb"

	hashMap := make(map[rune]int)

	for _, c := range s {

		if _, ok := hashMap[c]; ok {
			hashMap[c] += 1
		} else {
			hashMap[c] = 1
		}

	}
	keys := maps.Keys(hashMap)
	fmt.Println(string(keys))
	fmt.Println()
}
