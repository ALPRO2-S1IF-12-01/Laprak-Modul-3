package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a < c || b < d {
		fmt.Println("Error: n harus lebih besar atau sama dengan r")
		return
	}
	fmt.Printf("P(%d,%d) = %d\n", a, c, permutasi(a, c))
	fmt.Printf("C(%d,%d) = %d\n", a, c, kombinasi(a, c))
	fmt.Printf("P(%d,%d) = %d\n", b, d, permutasi(b, d))
	fmt.Printf("C(%d,%d) = %d\n", b, d, kombinasi(b, d))
}

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
