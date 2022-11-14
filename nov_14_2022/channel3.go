package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello world channels")

}

func main() {
	go hello()
	time.Sleep(1 * time.Millisecond)
	fmt.Println("main function")
}
