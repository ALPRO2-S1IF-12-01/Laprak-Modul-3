// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func compose(f, g, h func(int) int, x int) int {
	return f(g(h(x)))
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(compose(f, g, h, a))
	fmt.Println(compose(g, h, f, b))
	fmt.Println(compose(h, f, g, c))
}
