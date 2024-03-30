package main

import (
	"sync"
	"testing"
)

var x = 0

func TestIncrement(t *testing.T) {
	var wg sync.WaitGroup
	// var m sync.Mutex
	expected := 20
	var actual int

	for i := 0; i < expected; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			x += 1

		}()
	}
	wg.Wait()
	actual = x

	if actual != expected {
		t.Errorf("Expected x to be %d, but got %d", expected, actual)

	}
}
