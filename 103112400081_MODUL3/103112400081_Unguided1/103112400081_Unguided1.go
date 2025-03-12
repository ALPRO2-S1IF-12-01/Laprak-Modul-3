// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	perm_ac := permutasi(a, c)
	komb_ac := kombiinasiEfisien(a, c)
	perm_bd := permutasi(b, d)
	komb_bd := kombiinasiEfisien(b, d)

	fmt.Println(perm_ac, komb_ac)
	fmt.Println(perm_bd, komb_bd)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombiinasiEfisien(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	if k > n/2 {
		k = n - k
	}
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}
