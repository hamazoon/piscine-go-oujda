package piscine

func ReverseMenuIndex(menu []string) []string {
	len := len(menu)
	rev := make([]string, len)
	for i := 0; i < len; i++ {
		rev[i] = menu[len-i-1]
	}
	return rev
}
