// Felix Pedrosa V

package main

import "fmt"

// Fungsi untuk menghitung faktorial
func hitungFaktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func hitungPermutasi(n, r int) int {
	return hitungFaktorial(n) / hitungFaktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func hitungKombinasi(n, r int) int {
	return hitungFaktorial(n) / (hitungFaktorial(r) * hitungFaktorial(n-r))
}

func main() {
	var x, y, z, w int
	fmt.Scan(&x, &y, &z, &w)

	if x >= z && y >= w {
		fmt.Println(hitungPermutasi(x, z), hitungKombinasi(x, z))
		fmt.Println(hitungPermutasi(y, w), hitungKombinasi(y, w))
	} else {
		fmt.Println("Input tidak valid, pastikan x >= z dan y >= w")
	}
}