package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
	}()

	for {
		select {
		case <-ch1:
			fmt.Println(<-ch1)
		default:
			fmt.Println("nothing")
		}

	}
}
