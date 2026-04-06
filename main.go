package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

var TEXT string = "Shivanshu"

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	for {
		fmt.Print("\033[2J")
		fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(TEXT)/2)
		fmt.Print(TEXT)
		time.Sleep(time.Second / 24)
	}
}
