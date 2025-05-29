package piscine

func IsAlpha(s string) bool {
	for _, i := range []rune(s) {
		if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') && (i < '1' || i > '9') {
			return false
		}
	}
	return true
}
