package main

import (
	"os"

	"golang.org/x/term"
)

var TEXT string = "Shivanshu"

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	print(width, height)
}
