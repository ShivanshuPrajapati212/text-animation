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

var TEXT string = "shivanshu"

func main() {
	fig := figure.NewFigure(TEXT, "", true)
	lines := fig.Slicify()

	longest := slices.MaxFunc(lines, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	})

	fmt.Print("\033[?25l")

	lllll := 0
	for {
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			panic(err)
		}
		fmt.Print("\033[2J")

		for i, v := range lines {
			fmt.Printf("\033[%d;%dH", ((height/2)-len(lines)/2)+i, (width/2)-len(longest)/2)
			for _, letter := range v {
				fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", lllll, lllll, lllll, string(letter))
				if lllll >= 255 {
					lllll = 0
				} else {
					lllll += 1
				}
			}
		}

		time.Sleep(time.Second / 24)
	}
}
