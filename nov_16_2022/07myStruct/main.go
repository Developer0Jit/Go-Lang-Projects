package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Struct in golang")
	// no inheritance in golang: No super parents
	jitendra := User{"jitendra", "jitu@gmail.com", true, 23}

	fmt.Println(jitendra)
	fmt.Printf("hitesh Details are: %+v\n", jitendra)
	fmt.Printf("Name is %v and Emai is %v", jitendra.Name, jitendra.Email)
}
