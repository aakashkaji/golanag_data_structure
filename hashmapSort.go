package main

import (
	"fmt"
	"sort"
)

func sortByValue(m map[string]int) []string {
	// Create a slice to hold the keys
	keys := make([]string, 0, len(m))
	fmt.Println(keys)

	// Append keys to the slice
	for k := range m {
		keys = append(keys, k)
	}

	// Sort the keys based on the corresponding values
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}

func main() {
	// Example map
	myMap := map[string]int{
		"apple":  3,
		"banana": 2,
		"orange": 5,
		"grape":  1,
	}

	// Sort the map by value
	sortedKeys := sortByValue(myMap)

	// Print the keys in sorted order
	for _, key := range sortedKeys {
		fmt.Printf("%s: %d\n", key, myMap[key])
	}
}
