package main

import "fmt"

func main() {
	var x [256]string

	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", i, v, []byte(v))
	}
}
