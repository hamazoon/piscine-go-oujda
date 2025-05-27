package piscine

func Sqrt(nb int) int {
	sqr := 1
	if nb <= 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			if i*i == nb {
				sqr = i
				return sqr
			}
		}
	}
	return 0
}
