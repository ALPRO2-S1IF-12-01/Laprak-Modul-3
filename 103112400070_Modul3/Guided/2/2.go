package main

import (
	"fmt"
)

func celciusToFarenheit(celcius float64) float64 {
	return (9.0/5.0)*celcius + 32

}
func main() {
	var N int
	fmt.Print("Masukkan jumlah data: ")
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid, pastikan memasukkan angka.")
		return

	}
	temperature := make([]float64, N)

	fmt.Println("Masukkan suhu dalam celcius: ")
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&temperature[i])
		if err != nil {
			fmt.Println("Input tidak valid, pastikan memasukkan angka positif")
			return
		}
	}
	fmt.Println("Suhu dalam farenheit: ")
	for _, temp := range temperature {
		fmt.Printf("%.2f\n", celciusToFarenheit(temp))

	}

}
