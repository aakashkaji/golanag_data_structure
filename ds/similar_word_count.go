package main

import (
	"fmt"
)

func main() {
	// Sample strings
	str1 := "hello"
	str2 := "world"

	// Initialize maps to store letter counts
	counts1 := make(map[rune]int)
	counts2 := make(map[rune]int)

	fmt.Println(counts1)

	// Count the occurrences of each letter in the first string
	for _, char := range str1 {
		counts1[char]++
	}

	fmt.Println(counts1)


	// Count the occurrences of each letter in the second string
	for _, char := range str2 {
		counts2[char]++
	}

	// Find the common letters between the two strings
	commonLetters := make([]rune, 0)
	for char := range counts1 {
		if counts2[char] > 0 {
			commonLetters = append(commonLetters, char)
		}
	}

	// Display the common letters
	fmt.Println("Common Letters:", string(commonLetters))
}
