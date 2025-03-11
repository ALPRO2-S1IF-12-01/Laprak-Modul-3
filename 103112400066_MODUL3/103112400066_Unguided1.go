// DWI OKTA SURYANINGRUM
package main

import "fmt"

func main() {
	// deklarasi variabel a, b, c, dan d bertipe integer
	var a, b, c, d int

	// menerima input untuk variabel a, b, c, dan d secara bersamaan
	fmt.Scan(&a, &b, &c, &d)

	// memeriksa apakah a >= c dan b >= d
	if a >= c && b >= d {
		// jika kondisi terpenuhi, menampilkan hasil permutasi dan kombinasi
		// dari (a, c) dan (b, d)
		fmt.Println(permutasi(a,c), kombinasi(a,c))
		fmt.Println(permutasi(b,d), kombinasi(b,d))
	} else {
		// Jika kondisi tidak terpenuhi, menampilkan pesan kesalahan
		fmt.Println("input tidak sesuai")
	}
}

// fungsi untuk menghitung faktorial dari n
func faktorial(n int) int {
	hasil := 1
	// menghitung hasil faktorial dengan perkalian berulang dari 1 hingga n
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	// mengembalikan hasil faktorial
	return hasil
}

// fungsi untuk menghitung permutasi, rumus: n! / (n-r)!
func permutasi(n, r int) int {
	// jika r lebih besar dari n, permutasi tidak valid, mengembalikan 0
	if r > n {
		return 0
	}
	// menghitung permutasi dengan menggunakan fungsi faktorial
	return faktorial(n) / faktorial(n-r)
}

// fungsi untuk menghitung kombinasi, rumus: n! / (r! * (n-r)!)
func kombinasi(n, r int) int {
	// jika r lebih besar atau sama dengan n, kombinasi tidak valid
	if r > n {
		return 0
	}
	// menghitung kombinasi dengan menggunakan fungsi faktorial
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
