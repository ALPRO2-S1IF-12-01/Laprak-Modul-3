package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fogoh := f(g(h(a)))
	gohof := g(h(f(b)))
	hofog := h(f(g(c)))

	fmt.Println(fogoh)
	fmt.Println(gohof)
	fmt.Println(hofog)

}

func f(n int) int {
	return n * n
}

func g(n int) int {
	return n - 2
}

func h(n int) int {
	return n + 1
}
