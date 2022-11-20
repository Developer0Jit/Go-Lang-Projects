package main

import "fmt"

func main() {
	fmt.Println("Welcome to video on slice")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

}
