package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	// fmt.Printf("%T %v \n", myAge, myAge)
	fmt.Printf("%T", myAge)
	fmt.Println()
	fmt.Printf("%v", myAge)
}
