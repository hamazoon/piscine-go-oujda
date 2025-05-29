package piscine

func ToUpper(s string) string {
	var str string
	for _, i := range []rune(s) {
		if i >= 'a' && i <= 'z' {
			i -= 32
		}
		str = str + string(i)
	}
	return str
}
