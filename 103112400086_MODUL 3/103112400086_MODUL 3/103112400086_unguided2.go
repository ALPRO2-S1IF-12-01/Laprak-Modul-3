package main

import "fmt"

var a, b, c int

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogog(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	fmt.Scan(&a, &b, &c)

	fmt.Printf("fogog(%d) = %d\n", a, fogog(a))
	fmt.Printf("gohof(%d) = %d\n", b, gohof(b))
	fmt.Printf("hofog(%d) = %d\n", c, hofog(c))

}
