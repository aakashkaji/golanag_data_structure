package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		fmt.Println("hi aakash")
	}()

	ch <- "naveen"
	ch <- "paul"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(2000)
}
