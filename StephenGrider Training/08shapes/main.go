package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sidelenght float64
}
type shape interface {
	getArea() float64
}

func main() {
	t := triangle{base:10, height:10}
	s := square{sidelenght: 10}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea()float64 {
	return 0.5 * t.base * t.height

}

func (s square) getArea()float64 {
	return s.sidelenght * s.sidelenght

}

func printArea(s shape) { 
	fmt.Println(s.getArea())
 }