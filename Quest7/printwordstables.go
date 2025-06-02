package piscine

import "github.com/01-edu/z01"

func splitWhiteSpaces(s string) []string {
	r := []string{}
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if s[index] != ' ' {
				r = append(r, s[index:i])
			}
			index = i + 1
		} else if i == len(s)-1 {
			r = append(r, s[index:])
		}
	}
	return r
}

func printstring(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	for _, str := range a {
		printstring(str)
	}
}
