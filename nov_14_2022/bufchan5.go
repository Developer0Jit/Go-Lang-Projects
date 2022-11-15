package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 5
	ch <- 4
	close(ch)
	for n := range ch {
		fmt.Println("Received", n)
	}
}
