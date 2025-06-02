package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a >= 1; a-- {
		for b := a - 1; b >= 0; b-- {
			printTwoDigit(a)
			z01.PrintRune(' ')
			printTwoDigit(b)
			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigit(n int) {
	z01.PrintRune(rune(n/10) + '0')
	z01.PrintRune(rune(n%10) + '0')
}
