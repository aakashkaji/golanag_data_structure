package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// try to write/explain generic programm

type num int32

func genericAdd[T constraints.Ordered](a, b T) T {
	return a + b
}


func simpleAdd(a, b int) int {
	return a + b
}

func main() {
	a, b := 10, 20

	result := simpleAdd(a, b)

	fmt.Println(result)

	k := num(4)

	fmt.Println(genericAdd(float32(k), 2.21))
}
