package main

import "fmt"

type purchase interface {
	sell()
}

type display interface {
	show()
}

type saleman interface {
	purchase
	display
}

type game struct {
	name, price    string
	gameCollection []string
}

func (t game) sell() {
	fmt.Println("----------------------------")
	fmt.Println("Name:", t.name)
	fmt.Println("Price", t.price)
	fmt.Println("---------------------------")
}

func (t game) show() {
	fmt.Println("The games availabe are")
	for _, name := range t.gameCollection {
		fmt.Println(name)
	}
	fmt.Println("--------------------------------------")

}

func shop(s saleman) {
	fmt.Println(s)
	s.sell()
	s.show()
}

func main() {

	collection := []string{"XYZ",
		"Trial by Code", "Sea of Rubies"}
	game1 := game{"ABC", "$125", collection}
	shop(game1)

}
