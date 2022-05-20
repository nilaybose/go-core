package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
