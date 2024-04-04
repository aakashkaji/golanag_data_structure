package main

import "fmt"

func main() {

	v := 15 % 4

	switch v {
	case 3:
		fmt.Println(100)
		fallthrough
	case 2:
		fmt.Println(42)
	case 1:
		fmt.Println(1000)
	}

}
