package piscine

func Rot14(s string) string {
	res := []rune{}
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			i = 'a' + ((i - 'a' + 14) % 26)
		} else if i >= 'A' && i <= 'Z' {
			i = 'A' + ((i - 'A' + 14) % 26)
		}
		res = append(res, i)
	}
	return string((res))
}
