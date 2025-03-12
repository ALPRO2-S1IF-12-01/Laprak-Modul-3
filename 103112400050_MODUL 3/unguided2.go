package main

import "fmt"

//ARIEL AHNAF KUSUMA 103112400050
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}

func f(n int) int { return n * n }

func g(n int) int { return n - 2 }

func h(n int) int { return n + 1 }
