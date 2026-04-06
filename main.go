package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/common-nighthawk/go-figure"
	"golang.org/x/term"
)

var TEXT string = "Shivanshu"

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	fig := figure.NewFigure(TEXT, "block", true)
	lines := fig.Slicify()

	longest := slices.MaxFunc(lines, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	})

	for {
		fmt.Print("\033[2J")

		for i, v := range lines {
			fmt.Printf("\033[%d;%dH", ((height/2)-len(lines)/2)+i, (width/2)-len(longest)/2)
			fmt.Print(v)
		}

		time.Sleep(time.Second / 24)
	}
}
