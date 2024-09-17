package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type circle struct {
	radius float64
}

func (z circle) area() float64 {
	return 3.14 * z.radius * z.radius
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{30}
	// fmt.Printf("%T\n", s)
	// fmt.Println(s.area())
	// fmt.Println("--------------")
	info(s)
	info(c)
}
