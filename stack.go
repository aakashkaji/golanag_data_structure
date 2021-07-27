package main

import "fmt"

// stack concept based on ***********LIFO -> Last In First Out*******//

type stack struct {
	items []int
}

// push element on stack
func (s *stack) push(el int) {
	s.items = append(s.items, el)
}

// pop elemrnt from stack
func (s *stack) pop() int {
	l := len(s.items)
	removedItem := s.items[l-1]
	s.items = s.items[:l-1]
	return removedItem
}

func main() {
	myStack := stack{}
	myStack.push(1)
	myStack.push(2)
	myStack.push(4)
	myStack.push(5)
	myStack.push(8)

	fmt.Println("After push element on stack:", myStack.items)

	// pop element

	fmt.Println(myStack.pop())

	fmt.Println("After pop element from stack::", myStack.items)

}
