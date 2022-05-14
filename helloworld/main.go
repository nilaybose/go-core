package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ab struct{}
type bc struct{}

func (s *ab) greeting() {
	fmt.Println("AB -> greeting")
}

func (s *bc) greeting() {
	fmt.Println("BC -> greeting")
}

func main() {
	d := deck{}

	cards := []string{"spade", "clubs", "hearts", "diamond"}
	values := []string{"1", "2", "3", "4"}

	for _, card := range cards {
		for _, value := range values {
			d = append(d, card+" of "+value)
		}
	}
	d.print()

	_, b, _ := abc()

	fmt.Println([]byte(b))

	println(b)

	println()

	ioutil.WriteFile("decks.txt", []byte(strings.Join(d, ":")), 0755)

	nums := []int{}

	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
	}

	for _, n := range nums {
		if n%2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}

	ab := ab{}
	ab.greeting()

	bc := bc{}
	bc.greeting()
}

func greeting(b bot) {
	fmt.Println(b.greeting())
}
