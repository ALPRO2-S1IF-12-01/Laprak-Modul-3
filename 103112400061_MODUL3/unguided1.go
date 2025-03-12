package main

import "fmt"

func main() {
	var a, b, c, d int
	
	fmt.Scanln(&a, &b, &c, &d)

	fmt.Println(permutasi(a,c), kombinasi(a,c))
	fmt.Println(permutasi(b,d), kombinasi(b,d))
}

func permutasi(x, y int) int {
	var n, hasil int
	n = factorial(x)
	hasil = n/factorial(x-y)
	return hasil
}

func kombinasi(x, y int) int {
	var n, r, hasil int
	n = factorial(x)
	r = factorial(y)
	hasil = n/(r*(factorial(x-y)))
	return hasil
}

func factorial(n int) int{
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}