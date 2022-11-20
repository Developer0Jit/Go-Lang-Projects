package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class on pointer")

	// var one *int

	// fmt.Println("the pointer value is:", one)

	myNumber := 24

	var ptr = &myNumber

	fmt.Println("the value of pointer is:", ptr)
	fmt.Println("the value of pointer is:", *ptr)

	*ptr = *ptr + 2

	fmt.Println("the value og myNumber is:", myNumber)
}
