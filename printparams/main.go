package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:len(os.Args)]
	for _, i := range arg {
		for _, j := range i {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
