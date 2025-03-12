package main

// Muhammad Faris Rachmadi
// 103112400079

import "fmt"

func f(n int) int {
	return n * n
}

func g(n int) int {
	return n - 2
}

func h(n int) int {
	return n + 1
}

func fogoh(n int) int {
	return f(g(h(n)))
}

func gohof(n int) int {
	return g(h(f(n)))
}

func hofog(n int) int {
	return h(f(g(n)))
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}
