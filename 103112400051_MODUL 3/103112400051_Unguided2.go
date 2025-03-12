//PRATAMA BINTANG DANISWARA 103112400051
package main
import "fmt"
func f(x int) int {
	hasil := x
	hasil *= x
	return hasil
}
func g(x int) int {
	hasil := x
	hasil -= 2
	return hasil
}
func h(x int) int {
	hasil := x
	hasil += 1
	return hasil
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
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}