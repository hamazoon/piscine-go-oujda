package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 33
	ptr.y = 34
}

func itoi(n int) string {
	if n == 0 {
		return "0"
	}
	sing := false
	if n < 0 {
		sing = true
		n = -n
	}
	str := ""
	for n > 0 {
		str = string(rune(n%10+'0')) + str
		n = n / 10
	}
	if sing {
		str = "-" + str
	}
	return str
}

func main() {
	points := &point{}

	setPoint(points)

	output := "x = " + itoi(points.x) + " , y = " + itoi(points.y) + "\n"
	for _, r := range output {
		z01.PrintRune(r)
	}
}
