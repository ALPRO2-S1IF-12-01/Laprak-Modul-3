// Anastasia Adinda Narendra Indrianto
// 103112400085 S1IF-12-01
package main

import "fmt"

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n <= 1 {
		return 1
	}
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

// Fungsi untuk menghitung permutasi
func hitungPermutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func hitungKombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var x, y, p, q int

	// Membaca input
	fmt.Scan(&x, &y, &p, &q)

	// Validasi input
	if x >= p && y >= q {
		fmt.Println(hitungPermutasi(x, p), hitungKombinasi(x, p))
		fmt.Println(hitungPermutasi(y, q), hitungKombinasi(y, q))
	} else {
		fmt.Println("Input kurang tepat")
	}
}
