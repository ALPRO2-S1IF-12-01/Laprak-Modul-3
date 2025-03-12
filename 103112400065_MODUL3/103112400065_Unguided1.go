package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b), kombinasi(a, b))
	} else {
		fmt.Println(permutasi(b, a), kombinasi(b, a))
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

func kombinasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / (faktorial(n-r)*faktorial(r))
}
