package main

import (
	"fmt"
	"math"
)

func main() {
	a := divide(5, 0)
	fmt.Println(a)
}

func divide(dividend, divisor float64) (q float64) {
	defer func() {
		if r := recover(); r != nil {
			q = math.NaN()
		}
	}()
	if divisor == 0 {
		panic("Divide by zero")
	}
	return dividend / divisor
}
