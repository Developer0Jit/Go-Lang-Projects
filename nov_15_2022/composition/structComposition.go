// Golang program to store information
// about games in structs and display them
package main

import "fmt"

// We create a struct details to hold
// generic information about games
type details struct {
	genre       string
	genreRating string
	reviews     string
}

// We create a struct game to hold
// more specific information about
// a particular game
type game struct {
	name  string
	price string
	// We use composition through
	// embedding to add the
	// fields of the details
	// struct to the game struct
	details
}

// this is a method defined
// on the details struct
func (d details) showDetails() {
	fmt.Println("Genre:", d.genre)
	fmt.Println("Genre Rating:", d.genreRating)
	fmt.Println("Reviews:", d.reviews)
}

// this is a method defined
// on the game struct
// this method has access
// to showDetails() as well since
// the game struct is composed
// of the details struct
func (g game) show() {
	fmt.Println("Name: ", g.name)
	fmt.Println("Price:", g.price)
	g.showDetails()
}

func main() {

	// defining a struct
	// object of Type details
	action := details{"Action", "18+", "mostly positive"}

	// defining a struct
	// object of Type game
	newGame := game{"XYZ", "$125", action}

	newGame.show()
}
