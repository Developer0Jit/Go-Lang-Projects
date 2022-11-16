package main

import "fmt"

type Dog struct {
	Breed   string
	Weight  int16
	Friends []string
}

func (d Dog) ValReceiver() {
	d.Breed = "Terrier"
	d.Weight = 10
	d.Friends[0] = "Chris"
}

func main() {
	dog := Dog{Breed: "Dalmation", Weight: 40, Friends: []string{"Fred", "Amy"}}

	fmt.Printf("dog before ValReceiver: %+v\n", dog) // {Breed:Dalmation Weight:40 Friends:[Fred Amy]}
	dog.ValReceiver()
	fmt.Printf("dog after ValReceiver: %+v\n", dog) // {Breed:Dalmation Weight:40 Friends:[Chris Amy]}
}
