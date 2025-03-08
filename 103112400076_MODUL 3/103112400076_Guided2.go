package main

import "fmt"

// Fungsi untuk mengonversi Celcius ke Fahrenheit
func celciusToFahrenheit(celcius float64) float64 {
	return (9.0/5.0)*celcius + 32
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah data: ")
	_, err := fmt.Scan(&N)

	// Validasi input jumlah data
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid, pastikan masukkan angka positif")
		return
	}

	temperatures := make([]float64, N) // Array untuk menyimpan suhu dalam Celcius

	// Membaca suhu dalam Celcius
	fmt.Println("Masukkan suhu dalam Celcius:")
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&temperatures[i])
		if err != nil {
			fmt.Println("Input tidak valid, pastikan masukkan angka")
			return
		}
	}

	// Konversi dan output hasil
	fmt.Println("Suhu dalam Fahrenheit:")
	for _, temp := range temperatures {
		fmt.Printf("%.2f\n", celciusToFahrenheit(temp)) // Perbaikan format
	}
}
