// DWI OKTA SURYANINGRUM
package main

import "fmt"

// Fungsi untuk mengonversi suhu dari Celsius ke Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 { 
	// Fungsi ini menerima suhu dalam Celsius dan mengonversinya ke Fahrenheit menggunakan rumus
	return (9.0/5.0)*celsius + 32 
}

func main() {
	// Deklarasi variabel N untuk jumlah suhu yang akan dikonversi
	var N int 
	fmt.Print("Masukkan jumlah data: ") 
	
	// Membaca input dari pengguna untuk jumlah data suhu yang akan dimasukkan
	_, err := fmt.Scan(&N) 
	// Melakukan validasi input, jika N kurang dari atau sama dengan 0 atau terdapat error, maka tampilkan pesan error
	if err != nil || N <= 0 { 
		fmt.Println("Input tidak valid, pastikan memasukkan angka positif.") 
		return
	}

	// Membuat slice temperatures untuk menyimpan suhu dalam Celsius sebanyak N data
	temperatures := make([]float64, N) 

	// Membaca N buah suhu dalam Celsius dari pengguna
	fmt.Println("Masukkan suhu dalam Celsius:") 
	for i := 0; i < N; i++ { 
		// Melakukan perulangan sebanyak N kali untuk membaca setiap suhu
		_, err := fmt.Scan(&temperatures[i]) 
		// Jika input bukan angka, maka tampilkan pesan error dan hentikan program
		if err != nil { 
			fmt.Println("Input tidak valid, pastikan memasukkan angka.")
			return
		}
	}

	// Mengonversi suhu ke Fahrenheit dan mencetak hasilnya
	fmt.Println("Suhu dalam Fahrenheit:") 
	// Perulangan untuk membaca setiap elemen suhu dalam slice temperatures
	for _, temp := range temperatures { 
		// Mencetak hasil konversi suhu dengan dua angka di belakang koma
		fmt.Printf("%.2f\n", celsiusToFahrenheit(temp)) 
	}
}