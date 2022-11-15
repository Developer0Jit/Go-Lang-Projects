package main

import "fmt"

type VowelFinder interface {
	FindeVovels() []rune
}

type Mystring string

func (ms Mystring) FindeVovels() []rune {
	var Vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			Vowels = append(Vowels, rune)
		}
	}
	return Vowels
}

func main() {
	name := Mystring("john Hooks")
	var v VowelFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindeVovels())

}
