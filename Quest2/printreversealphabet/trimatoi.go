package piscine

func TrimAtoi(s string) int {

	if len(s) == 0 {
		return 0
	}
	isn := false
	fnd := false
	rus := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			rus = rus*10 + int(rune(s[i]-'0'))
			fnd = true
		} else if s[i] == '-' && fnd == false {
			isn = true
		}
	}
	if isn {
		return -rus
	}
	return rus
}
