package piscine

func IsPrintable(s string) bool {
	for _, i := range []rune(s) {
		if i < 32 || i > 126 {
			return false
		}
	}
	return true
}
