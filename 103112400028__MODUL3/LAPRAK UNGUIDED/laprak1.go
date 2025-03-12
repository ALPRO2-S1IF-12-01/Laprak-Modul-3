// MUHAMMAD GAMEL AL GHIFARI
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permu(a, c), kombi(a, c))
	fmt.Println(permu(b, d), kombi(b, d))
}

func fakt(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func permu(n, r int) int {
	return fakt(n) / fakt(n-r)
}
func kombi(n, r int) int {
	return fakt(n) / (fakt(r) * fakt(n-r))
}