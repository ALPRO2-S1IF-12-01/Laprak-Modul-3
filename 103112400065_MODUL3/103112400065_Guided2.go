package main
import "fmt"
func celciusToFahrenheit(celcius float64) float64 {
	return (9.0/5.0)*celcius + 32
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah data: ")
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid, pastikan masukkan angka positif")
		return
	} 
	tempereatures := make([]float64, N)

	// membaca suhu dalam celcius
	fmt.Println("Masukkan suhu dalam celcius:")
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&tempereatures[i])
		if err != nil {
			fmt.Println("Input tidak valid, pastikan masukkan angka")
			return
		}
	}

	//Mengonversi ke fahrenheit dan mencetak hasil
	fmt.Println("Suhu dalam fahrenheit:")
	for _, temp := range tempereatures {
		fmt.Printf("%.2f\n", celciusToFahrenheit(temp))
	}
}