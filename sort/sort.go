package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p := []Person{
		{"Nilay", 41},
		{"John", 48},
		{"David", 35},
	}

	sort.Sort(ByAge(p))
	fmt.Printf("%+v", p)
	fmt.Println()

	sort.Slice(p, func(i, j int) bool {
		return p[i].Age > p[j].Age
	})

	fmt.Printf("%+v", p)
	fmt.Println()
	sort.Sort(sort.Reverse(ByAge(p)))
	fmt.Printf("%+v", p)

	s := []int{5, 2, 6, 3, 1, 4} // unsorted

	// sort.Slice(s, func(i, j int) bool {
	// 	return s[i] > s[j]
	// })

	fmt.Println()
	fmt.Println(s)
}
