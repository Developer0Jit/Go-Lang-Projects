package main

import (
	"fmt"
	"time"
)

func main() {
	go greets("Hello")
	greets("Worlds")
}

func greets(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}

}
