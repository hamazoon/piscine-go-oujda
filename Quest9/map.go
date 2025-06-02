package piscine

func Map(f func(int) bool, a []int) []bool {
	booly := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			booly[i] = true
		} else {
			booly[i] = false
		}
	}
	return booly
}
