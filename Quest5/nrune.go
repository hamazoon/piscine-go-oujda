package piscine

func NRune(s string, n int) rune {
	nr := []rune(string(s))
	if n > 0 && len(nr) >= n {
		return nr[n-1]
	} else {
		return 0
	}
}
