package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output" + "\n"
	}
	cnt := 0
	var s string
	for i, char := range str {
		if char != ' ' && cnt != 5 {
			s += string(char)
			cnt++
		} else if cnt == 5 {
			s += " "
			cnt = 0
		}
		if i == len(str)-1 && len(s) > 0 && s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
		}
	}
	return s + "\n"
}
