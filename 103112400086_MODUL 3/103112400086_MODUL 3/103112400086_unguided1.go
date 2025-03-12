package main

import "fmt"

func main() {
	var a, b, c, d, p1, p2, c1, c2 int
	fmt.Scan(&a, &b, &c, &d)

	p1 = permutasi(a, c)
	c1 = combination(a, c)

	p2 = permutasi(b, d)
	c2 = combination(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}

func combination(n, r int) int {
	if r > n {
		return 0
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
