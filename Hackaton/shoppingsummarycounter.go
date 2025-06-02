package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	countwrd := make(map[string]int)
	count := 0
	ch := ""
	s := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			ch += string(str[i])
		} else {
			s = append(s, ch)
			ch = ""
		}
	}
	s = append(s, ch)
	for k := 0; k < len(s); k++ {
		for j := 0; j < len(s); j++ {
			if s[k] == s[j] {
				count++
			}
		}
		countwrd[s[k]] = count
		count = 0
	}
	return countwrd
}
