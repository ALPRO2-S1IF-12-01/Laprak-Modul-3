//Raihan Adi Arba
//103112400071

package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}

func factorial(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi permutasi: P(n, r) = n! / (n-r)!
func permutation(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

// Fungsi kombinasi: C(n, r) = n! / (r!(n-r)!)
func combination(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}
