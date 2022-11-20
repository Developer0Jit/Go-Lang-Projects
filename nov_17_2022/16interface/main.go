package main

import (
	"fmt"
)

type geomentry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height //area of rectangle r=w*h
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height //rectangle of perimeter r=2(w+h)
}

func measure(g geomentry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	fmt.Println("welcome to interface")

	r := rect{width: 3, height: 4}

	measure(r)
}
