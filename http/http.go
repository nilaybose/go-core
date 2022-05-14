package main

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	sidelength float64
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func getArea(s shape) float64 {
	return s.getArea()
}

// func main() {
// 	var t shape = triangle{base: 5, height: 6}
// 	var s shape = square{sidelength: 10}

// 	fmt.Println("Square area: ", getArea(s))
// 	fmt.Println("Triangle area: ", getArea(t))
// }
