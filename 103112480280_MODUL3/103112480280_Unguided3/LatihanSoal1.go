// Nama : Anggun Wahyu W. (103112480280)
package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func permutation(n, r int) int {
	if n < r {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	if n < r {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	permutasiA := permutation(a, c)
	kombinasiA := combination(a, c)
	permutasiB := permutation(b, d)
	kombinasiB := combination(b, d)

	fmt.Printf("%d %d\n", permutasiA, kombinasiA)
	fmt.Printf("%d %d\n", permutasiB, kombinasiB)
}
