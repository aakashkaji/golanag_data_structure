package main

import "fmt"

func recoverFun() {
	if rec := recover(); rec != nil {
		fmt.Println("Recover from error rec", rec)
	}

}

func fn(name string){
	defer recoverFun()

	if name == "" {
		panic("name must be present")
	}

}

func main() {


	name := ""
	fn(name)
	

	fmt.Println("Normal flow of main thread....")
	fmt.Println("Normal flow of main thread....")


}
