// Input: s = "()"
// Output: true
// Example 2:
// Input: s = "()[]{}"
// Output: true

package main

import (
	"fmt"
)

func isValidString(s string) bool {

	charStack := make([]rune, 0)

	// store bracket in hashmap

	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range s {
		if value, ok := brackets[char]; ok {
			// pushed to stack
			charStack = append(charStack, value)

		} else {

			if len(charStack) == 0 || char != charStack[len(charStack)-1] {
				return false
			}
			charStack = charStack[:len(charStack)-1]

		}
		fmt.Println(char)
	}

	return len(charStack) == 0
}

func main() {

	s := "({}[])[]{}"

	fmt.Println(s)
	fmt.Println(isValidString(s))

}
