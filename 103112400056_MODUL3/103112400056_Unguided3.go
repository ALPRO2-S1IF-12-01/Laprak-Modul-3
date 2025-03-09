// Felix Pedrosa V

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

// Fungsi untuk memeriksa apakah titik berada di dalam lingkaran
func titikDalamLingkaran(centerX, centerY, radius, pointX, pointY float64) bool {
	return hitungJarak(centerX, centerY, pointX, pointY) <= radius
}

func main() {
	var centerX1, centerY1, radius1 float64
	var centerX2, centerY2, radius2 float64
	var koordinatX, koordinatY float64

	fmt.Scan(&centerX1, &centerY1, &radius1)
	fmt.Scan(&centerX2, &centerY2, &radius2)
	fmt.Scan(&koordinatX, &koordinatY)

	// Memeriksa apakah titik berada di dalam lingkaran
	dalamLingkaran1 := titikDalamLingkaran(centerX1, centerY1, radius1, koordinatX, koordinatY)
	dalamLingkaran2 := titikDalamLingkaran(centerX2, centerY2, radius2, koordinatX, koordinatY)

	// Menampilkan hasil
	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}