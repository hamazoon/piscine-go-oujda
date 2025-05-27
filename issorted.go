package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	dec := true
	cr := true
	for i := 0; i+1 < len(a); i++ {
		if f(a[i], a[i+1]) > 0 {
			cr = false
		}
		if f(a[i+1], a[i]) > 0 {
			dec = false
		}
	}
	return dec || cr
}

func f(a, b int) int {
	return a - b
}
