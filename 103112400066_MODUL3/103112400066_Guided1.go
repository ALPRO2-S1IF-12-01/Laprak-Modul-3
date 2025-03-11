// DWI OKTA SURYANINGRUM
package main

import "fmt"

// Fungsi utama untuk menjalankan program
func main() {
	var a, b int // mendeklarasikan dua variabel bertipe integer untuk menyimpan input dari pengguna
	fmt.Scan(&a, &b) // membaca input dua angka yang diberikan oleh pengguna dan menyimpannya dalam variabel a dan b

	// mengecek apakah nilai a lebih besar atau sama dengan b
	if a >= b { 
		// jika a lebih besar atau sama dengan b, maka memanggil fungsi permutasi dengan argumen a dan b
		fmt.Println(permutasi(a, b)) 
	} else {
		// jika b lebih besar dari a, maka memanggil fungsi permutasi dengan argumen b dan a
		fmt.Println(permutasi(b, a)) 
	}
}

// Fungsi faktorial digunakan untuk menghitung faktorial dari suatu bilangan
func faktorial(n int) int {
	hasil := 1 // inisialisasi variabel hasil dengan nilai awal 1
	// perulangan untuk menghitung faktorial dari 1 hingga n
	for i := 1; i <= n; i++ { 
		hasil *= i // mengalikan nilai hasil dengan i pada setiap iterasi
	}
	// mengembalikan hasil faktorial
	return hasil
}

// fungsi permutasi digunakan untuk menghitung permutasi dari dua angka n dan r
func permutasi(n, r int) int {
	// memeriksa apakah r lebih besar dari n, jika iya permutasi tidak valid
	if r > n { 
		return 0 // mengembalikan 0 jika r lebih besar dari n karena permutasi tidak dapat dihitung
	}
	// jika r tidak lebih besar dari n, maka menghitung permutasi dengan rumus: n! / (n - r)!
	// pertama menghitung faktorial dari n, lalu membaginya dengan faktorial dari (n - r)
	return faktorial(n) / faktorial(n - r)
}

// n dalam kode berasal dari input yang dimasukkan user.
// Kalau a lebih besar atau sama dengan b, maka a akan dianggap sebagai n dan b sebagai r.
// Kalau b lebih besar dari a, maka b akan dianggap sebagai n dan a sebagai r.