package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}
func dalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	fmt.Println("Lingkaran 1:")
	fmt.Print("Pusat (x y): ")
	fmt.Scan(&cx1, &cy1)
	fmt.Print("Jari-jari: ")
	fmt.Scan(&r1)
	fmt.Println("\nLingkaran 2:")
	fmt.Print("Pusat (x y): ")
	fmt.Scan(&cx2, &cy2)
	fmt.Print("Jari-jari: ")
	fmt.Scan(&r2)
	fmt.Println("\nTitik yang akan dicek:")
	fmt.Print("Koordinat (x y): ")
	fmt.Scan(&x, &y)
	if r1 <= 0 || r2 <= 0 {
		fmt.Println("\nError: Jari-jari harus lebih besar dari 0!")
		return
	}
	dalamLingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
	dalamLingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)
	fmt.Println("\nHasil Pengecekan:")
	fmt.Printf("Titik (%.2f, %.2f) ", x, y)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("berada di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("berada di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("berada di dalam lingkaran 2")
	} else {
		fmt.Println("berada di luar lingkaran 1 dan 2")
	}
}
