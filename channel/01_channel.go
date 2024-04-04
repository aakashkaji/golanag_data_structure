package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)

	return rand.Intn(1000)
}

func main() {
	numChan := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				numChan <- doWork()

			}()

		}
		wg.Wait()
		close(numChan)
	}()

	// retrive data from channels

	// for v := range numChan {
	// 	fmt.Println(v)
	// }

	for {
		select {
		case r, ok := <-numChan:
			if !ok {
				fmt.Println("-----------")
				return
			}
			fmt.Println(r, ok)
		default:
			fmt.Println("Waiting........")
			// wait for milliseconds
			time.Sleep(time.Millisecond)

		}
	}

	fmt.Println("aakash")

}
