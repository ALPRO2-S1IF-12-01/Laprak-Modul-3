package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)
	fmt.Println("\nHasil Komposisi Fungsi:")
	fmt.Printf("f(g(h(%d))) = %d\n", a, fogoh(a))
	fmt.Printf("g(h(f(%d))) = %d\n", b, gohof(b))
	fmt.Printf("h(f(g(%d))) = %d\n", c, hofog(c))
}
func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func fogoh(x int) int {
	return f(g(h(x)))
}
func gohof(x int) int {
	return g(h(f(x)))
}
func hofog(x int) int {
	return h(f(g(x)))
}
