package main

import (
	"fmt"
	"time"
)

func process(chss chan string) {
	time.Sleep(10500 * time.Millisecond)
	chss <- "process successful"
}

func main() {
	strChan := make(chan string)

	go process(strChan)

	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-strChan:
			fmt.Println(v)
			return
		default:
			fmt.Println("No value Received...")
		}
	}

}
