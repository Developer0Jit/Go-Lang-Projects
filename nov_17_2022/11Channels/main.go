package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1

	b, ok := <-ch

	fmt.Println(b, ok)

	close(ch)

	c, ok := <-ch

	fmt.Println(c, ok)
}
