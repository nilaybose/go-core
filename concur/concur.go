package main

import (
	"fmt"
	"runtime"
	"sync"
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
	sg.Add(100)

	var mu sync.Mutex

	for i := 0; i < 100; i++ {
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
}
