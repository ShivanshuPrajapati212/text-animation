package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"time"

	"github.com/common-nighthawk/go-figure"
	"golang.org/x/term"
)

var TEXT string = "Hello, Bot"

func main() {
	fig := figure.NewFigure(TEXT, "", true)
	lines := fig.Slicify()

	longest := slices.MaxFunc(lines, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	})

	fmt.Print("\033[?25l")

	for {
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			panic(err)
		}
		fmt.Print("\033[2J")

		for i, v := range lines {
			fmt.Printf("\033[%d;%dH", ((height/2)-len(lines)/2)+i, (width/2)-len(longest)/2)
			for _, letter := range v {
				fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", rand.Intn(256), rand.Intn(256), rand.Intn(256), string(letter))
			}
		}

		time.Sleep(time.Second / 24)
	}
}
