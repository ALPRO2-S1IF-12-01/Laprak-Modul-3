package main
//NUFAIL ALAUDDIN TSAQIF
//103112400084
import "fmt"

func kuadrat(n int) int {
	return n * n
}
func kurang(n int) int {
	return n - 2
}
func tambah(n int) int {
	return n + 1
}
func fogoh(n int) int {
	return kuadrat(kurang(tambah(n)))
}
func gohof(n int) int {
	return kurang(tambah(kuadrat(n)))
}
func hofog(n int) int {
	return tambah(kuadrat(kurang(n)))
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}
