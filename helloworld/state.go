package main

import "fmt"

type deck []string

type bot interface {
	greeting() string
	greeting1() string
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func abc() (string, string, string) {
	return "a", "b", "c"
}
