package main

import "fmt"

// Use Decorator pattern to get full name.

type full_name func(last_name string) string

func get_name(first_name string) func(last_name string) string {

	var f full_name = func(last_name string) string {

		return fmt.Sprintf("%s %s", first_name, last_name)
	}

	return f
}

func main() {

	var first_name, last_name string = "Aakash", "Gupta"

	// call get_name method which return func

	f := get_name(first_name)

	fmt.Println(f(last_name))



}