// package main

// import (
// 	"sync"
// 	"time"
// )

// func main() {
// 	c := make(chan string)

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		c <- `foo`
// 		c <- `bar`
// 	}()

// 	go func() {
// 		defer wg.Done()

// 		time.Sleep(time.Second * 1)
// 		println(`Message:` + <-c)
// 		println(`Message:` + <-c)
// 	}()

// 	wg.Wait()
// }

package main

import (
	"fmt"
	"time"
)

func access(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("start accessing channel")

	for i := range ch {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	// only modify this line to defined the capacity
	ch := make(chan int, 3)
	defer close(ch)

	go access(ch)

	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Filled")
	}

	time.Sleep(3 * time.Second)
}
