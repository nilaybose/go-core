package utils

import "fmt"

func Info() {
	fmt.Println("Utilities to be imported")
}

func concat(strs []string) string {
	var s string = ""
	for _, str := range strs {
		s += str + " "
	}
	return s
}
