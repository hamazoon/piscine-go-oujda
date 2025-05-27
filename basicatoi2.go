package piscine

func BasicAtoi2(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if int(s[i]) < '0' || int(s[i]) > '9' {
			return 0
		} else {
			n = n*10 + int(s[i]) - '0'
		}
	}
	return n
}
