package main

import "fmt"

type Perion struct {
	Name string
	Age  int
}

func ValueReceiver(p Perion) {
	p.Name = "John"
	fmt.Println("Inside value Recediver", p.Name)
}

func PointerReceiver(p *Perion) {
	p.Age = 24
	fmt.Println("Inside Pointer Receiver", p.Age)
}

func main() {
	p := Perion{"Tome", 24}
	p1 := &Perion{"jonny", 68}
	ValueReceiver(p)
	fmt.Println("Inside Main After Value Receiver", p.Name)
	PointerReceiver(p1)
	fmt.Println("Inside main after pointer Receiver", p1.Age)

}
