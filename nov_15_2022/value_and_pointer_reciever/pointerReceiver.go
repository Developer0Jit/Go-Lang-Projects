package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int16
}

func main() {
	dog := Dog{Breed: "Dalmation", Weight: 40}

	fmt.Printf("dog before PtrReceiver %+v\n", dog) // {Breed:Dalmation Weight:40}
	dog.PtrReceiver()
	fmt.Printf("dog after PtrReceiver: %+v\n", dog) // {Breed:Terrier Weight:10}
}
func (d *Dog) PtrReceiver() {
	d.Breed = "Terrier"
	d.Weight = 10
}
