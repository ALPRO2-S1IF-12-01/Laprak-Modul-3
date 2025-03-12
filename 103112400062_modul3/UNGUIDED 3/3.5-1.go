// M. DAVI ILYAS RENALDO
// 103112400062
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	perm_ac := permutasi(a, c)
	komb_ac := kombinasi(a, c)
	perm_bd := permutasi(b, d)
	komb_bd := kombinasi(b, d)

	fmt.Println(perm_ac, komb_ac)
	fmt.Println(perm_bd, komb_bd)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
