package piscine

func JumpOver(str string) string {
	var res string

	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}

	return res + "\n"
}
