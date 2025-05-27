package piscine

func IterativeFactorial(nb int) int {
	fac := 1
	if nb < 0 || nb > 21 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb == 1 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		fac = fac * i
	}
	return fac
}
