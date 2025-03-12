package main

import "fmt"

func main() {
	var to, lol int
	fmt.Scan(&to, &lol)

	if to >= lol {
		fmt.Println(permutasi(to, lol))
	} else {
		fmt.Println(permutasi(lol, to))
	}
}

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
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
