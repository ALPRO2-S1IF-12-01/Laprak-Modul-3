//Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func main() {
	var w, x, y, z int
	fmt.Scan(&w, &x, &y, &z)
	fmt.Println(mutasi(w, y), kombinasi(w, y))
	fmt.Println(mutasi(x, z), kombinasi(x, z))
}

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func mutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}