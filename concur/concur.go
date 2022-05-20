package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/nilay.bose/go-core/utils"
)

func main() {
	fmt.Println("Demo Race Condition")

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	var counter int = 0
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	var sg sync.WaitGroup
	sg.Add(2)

	var mu sync.Mutex

	utils.Info()

	for i := 0; i < 2; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			sg.Done()
		}()
	}

	sg.Wait()

	fmt.Println("Counter = ", counter)

	odd := make(chan int)
	even := make(chan int)
	end := make(chan int)

	go func(odd, even chan<- int) {
		type hotdog int
		var x hotdog = 9
		fmt.Printf("%T\n", x)

		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(even)
		close(odd)
	}(odd, even)

	go func() {
		var wg sync.WaitGroup
		wg.Add(2)

		go func(odd <-chan int, end chan<- int) {
			for i := range odd {
				end <- i
			}
			wg.Done()
		}(odd, end)

		go func(even <-chan int, end chan<- int) {
			for i := range even {
				end <- i
			}
			wg.Done()
		}(even, end)

		wg.Wait()
		close(end)
	}()

	for n := range end {
		fmt.Print(n, " ")
	}

}
