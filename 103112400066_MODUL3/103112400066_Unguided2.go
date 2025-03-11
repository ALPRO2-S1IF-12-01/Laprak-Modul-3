// DWI OKTA SURYANINGRUM
package main

import "fmt"

func main() {
	// deklarasi variabel a, b, dan c bertipe integer
	var a, b, c int 
	
	// membaca inputan dari user dan menyimpannya di variabel a, b, dan c
	fmt.Scan(&a, &b, &c)
	
	// memanggil fungsi fx, gx, dan hx secara berurutan dengan parameter yang sesuai
	// melakukan perhitungan fx(gx(hx(a))) dan menampilkan hasilnya
	fmt.Println(fx(gx(hx(a)))) 
	
	// melakukan perhitungan gx(hx(fx(b))) dan menampilkan hasilnya
	fmt.Println(gx(hx(fx(b)))) 
	
	// melakukan perhitungan hx(fx(gx(c))) dan menampilkan hasilnya
	fmt.Println(hx(fx(gx(c)))) 
}

// fungsi untuk menghitung x^2
func fx(x int) int {
	// mengembalikan hasil perkalian x dengan dirinya sendiri (x^2)
	return x * x 
}

// fungsi untuk menghitung x - 2
func gx(x int) int {
	// mengembalikan hasil pengurangan x dengan 2
	return x - 2 
}

// fungsi untuk menghitung x + 1
func hx(x int) int {
	// mengembalikan hasil penjumlahan x dengan 1
	return x + 1 
}
