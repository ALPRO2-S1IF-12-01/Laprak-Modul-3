package main

import "fmt"

//ARIEL AHNAF KUSUMA 103112400050

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int { return factorial(n) / factorial(n-r) }
func combination(n, r int) int { return factorial(n) / (factorial(r) * factorial(n-r)) }

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c), permutation(b, d), combination(b, d))
}
