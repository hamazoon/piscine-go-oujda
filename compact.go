package piscine

func Compact(ptr *[]string) int {
	arr := *ptr
	compact := []string{}
	for _, str := range arr {
		if str != "" {
			compact = append(compact, str)
		}
	}
	*ptr = compact
	return len(compact)
}
