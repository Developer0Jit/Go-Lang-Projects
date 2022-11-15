package main

import (
	"fmt"
)

type geomenty interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func measure(g geomenty) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	r := rect{width: 6, height: 7}

	measure(r)

}
