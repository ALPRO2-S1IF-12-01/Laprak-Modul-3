//Raihan Adi Arba
//103112400071

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	n, r := a, b
	if a < b {
		n, r = b, a
	}

	fmt.Println(permutasi(n, r), kombinasi(n, r))
}

func faktorial(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / (faktorial(n-r) * faktorial(r))
}
