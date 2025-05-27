package piscine

func StrLen(s string) int {
	var a int
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		a = a + 1
	}
	return a
}
