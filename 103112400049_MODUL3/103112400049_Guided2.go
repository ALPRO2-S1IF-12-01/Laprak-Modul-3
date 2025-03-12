package main

import "fmt"

func celciustofahrenheit(celcius float64) float64 {
	return (9.0/5.0)*celcius + 32
}

func main() {
	var n int
	fmt.Print("Masukan jumlah data: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Input tidak valid, pastikan memasukan input positif.")
		return
	}
	temperature := make([]float64, n)
	fmt.Println("Mauskan suhu dalam Celcius: ")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&temperature[i])
		if err != nil {
			fmt.Println("Input tidak valid, pastikan memasukan input angka.")
			return
		}
	}
	fmt.Println("Suhu dalam fahrenheit")
	for _, temp := range temperature {
		fmt.Printf("%.2f\n", celciustofahrenheit(temp))
	}
}
