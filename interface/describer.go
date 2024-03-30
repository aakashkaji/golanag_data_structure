package main

import "fmt"

type Human interface {
	Speak() string
	Walk() int
}

type Man struct {
	Language string
	Age      int
}

func (m *Man) Speak() string {
	return m.Language
}

func (m Man) Walk() int {
	return m.Age * 5

}

func main() {

	var h Human
	// var man Man
	man := Man{Language: "english", Age: 20}
	h = &man

	fmt.Println(h.Speak())
	fmt.Println(h.Walk())

}
