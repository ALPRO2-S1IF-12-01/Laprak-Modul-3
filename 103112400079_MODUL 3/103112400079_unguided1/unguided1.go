package main

// Muhammad Faris Rachmadi
// 103112400079

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	r := 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}

func permu(n, r int) int {
	return fact(n) / fact(n-r)
}

func combi(n, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("a b c d: ")
	fmt.Scan(&a, &b, &c, &d)

	if a < c || b < d {
		fmt.Println("Syarat salah")
		return
	}

	fmt.Println(permu(a, c), combi(a, c))
	fmt.Print(permu(b, d), combi(b, d))
}
