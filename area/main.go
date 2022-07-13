package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	sq := square{sideLength: 4.0}
	tr := triangle{base: 5.0, height: 3.0}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println("The area is: ", s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
