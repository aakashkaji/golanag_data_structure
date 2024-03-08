package main

import (
	"fmt"
	"time"
)

func recoverError() {

	if r := recover(); r != nil {
		fmt.Println("Recover code from", r)
	}
	
}

func returnsElement(sl [4]int) {

	defer recoverError()
	fmt.Println(sl[1])
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("Return func")
}

func main() {

	arr := [4]int{1, 2, 3, 4}

	go returnsElement(arr)
	time.Sleep(2000)
	fmt.Println("main func executed.")
}
