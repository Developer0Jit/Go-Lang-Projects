package main

import "fmt"

type author struct {
	name      string
	branch    string
	particles int
}

func (a *author) show(abranch string) {
	(*a).branch = abranch
}

func main() {
	res := author{
		name:   "john",
		branch: "IT",
	}
	fmt.Println("Authors name", res.name)
	fmt.Println("Authors name", res.branch)

	p := &res

	p.show("cse")
	fmt.Println("Authors name", res.name)
	fmt.Println("Authors name", res.branch)
}
