package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	n, open := <-ch
	fmt.Println("Received:%d, open:%t", n, open)
	n, open = <-ch
	fmt.Println("Received:%d, open:%t", n, open)
	n, open = <-ch
	fmt.Println("Received:%d, open:%t", n, open)

}
