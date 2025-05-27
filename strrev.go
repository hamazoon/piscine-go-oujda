package piscine

func StrRev(s string) string {
	var l string
	for _, v := range s {
		l = string(v) + l
	}
	return l
}
