package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	fogoh := f(g(h(a)))
	gohof := g(h(f(b)))
	hofog := h(f(g(c)))
	fmt.Printf("%v\n%v\n%v\n", fogoh, gohof, hofog)
}

func f(x int) int{
	return x*x
}

func g(x int) int{
	return x-2
}

func h(x int) int{
	return x+1
}