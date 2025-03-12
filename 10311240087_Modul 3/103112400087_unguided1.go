//103112400087 Muhammad Shabrian fadly

package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("masukkan nilai a, b, c, d : ")
	fmt.Scan(&a, &b, &c, &d)

	if a < c || b < d {
		fmt.Println("syarat tidak terpenuhi: a harus lebih besar atau sama dengan c, dan b lebih besar atau sama dengan d.")
		return
	}
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
