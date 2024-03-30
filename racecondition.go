package main

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {

	defer wg.Done()
	m.Lock()
	x += 1
	m.Unlock()

	fmt.Println("value of x==", x)

}
func main() {

	var wg sync.WaitGroup
	var m sync.Mutex
	fmt.Println("Started programm...")
	for i := 0; i < 20; i++ {

		wg.Add(1)
		go increment(&wg, &m)
	}
	wg.Wait()

}
