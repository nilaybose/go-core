package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Args[1]

	fmt.Println("Input File Name: ", f)

	file, err := os.Open(f)
	if err != nil {
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

}
