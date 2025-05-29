package piscine 

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	} 
	if n == 0 {
		z01.PrintRune('0')
	}

		var arr []int
		for   n>0 {
			arr = append(arr ,n % 10)
			n = n / 10  
		}
	
	for i := 0 ; i < len(arr) - 1 ; i++ {
		for j := 0 ; j < len(arr) - i - 1 ; j++ {
			if arr[j ] > arr[j + 1] {
				arr[j] , arr[j +1] = arr[j +1] , arr[j]
			}
		}
	}
	for n := 0 ; n < len(arr) ;n++ {
		z01.PrintRune(rune(arr[n] + '0'))
	}
}
