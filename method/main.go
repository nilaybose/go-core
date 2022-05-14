package main

import "fmt"

type Person struct {
	name string
}

type Action interface {
	speak()
}

func (p Person) speak() {
	fmt.Println(p.name)
}

func Actionable(act Action) {
	act.speak()
}

func main() {
	fmt.Println("Hello")

	p1 := Person{
		"hello p1",
	}

	p2 := &Person{
		"hello p2",
	}

	p1.speak()
	p2.speak()

	Actionable(p2)
	Actionable(&p1)
}
