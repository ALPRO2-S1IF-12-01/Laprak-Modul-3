package main 
import ("fmt")

func celciusToFahrenheit(celcius float64) float64 {
	return (9.0/5.0)*celcius + 32
}

func main(){
	var N int
	fmt.Print("Masukan jumlah data: ")
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid, pasti memasukan angka positif.")
		return
	}
	temperatures := make([]float64, N)

	fmt.Println("Masukan suhu dalam Celcius: ")
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&temperatures[i])
		if err != nil {
			fmt.Println("Input tidak valid, pastikan memasukan angka")
			return
		}
	}

	fmt.Println("Suhu dalam Fahrenheit:")
	for _, temp := range temperatures {
		fmt.Printf("%.2f\n", celciusToFahrenheit(temp))
	}
}