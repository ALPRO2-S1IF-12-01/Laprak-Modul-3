package main

import "fmt"

func main() { //main redeclared in this block
	var a, b int
	fmt.Scan(&a, &b)

	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int { //nama setiap fungsi harus berbeda", jika sama nanti error
	if r > n {
		return 0
	}
	return faktorial(n) / faktorial(n-r) //nilai n hrs lebih besar dr nilai r
}
