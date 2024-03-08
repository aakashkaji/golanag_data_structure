// number generator, square calculator, then value printer

package main

import "fmt"

func main() {

	number := make(chan int)
	square := make(chan int)

	go func() {

		defer close(number)
		for i := 1; i < 21; i++ {
			number <- i
		}
	}()

	// calculate square of number

	go func() {
		defer close(square)
		for x := range number {

			square <- x * x
		}
	}()

	for value := range square {

		fmt.Println("value of xquare:::", value)
	}
}
