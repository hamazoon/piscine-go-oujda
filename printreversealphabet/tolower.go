package piscine

func ToLower(s string) string {
	var str string
	for _, i := range []rune(s) {
		if i >= 'A' && i <= 'Z' {
			i += 32
		}
		str = str + string(i)
	}
	return str
}
