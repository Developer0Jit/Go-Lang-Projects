package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello world go routines")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("Main Function")
}
