// DWI OKTA SURYANINGRUM
package main

import (
	"fmt"
	"math"
)

func main() {
	// deklarasi variabel untuk pusat dan jari-jari lingkaran 1 dan 2, serta titik sembarang
	var cy1, cx1, r1 int
	var cy2, cx2, r2 int
	var y, x int

	// membaca input untuk lingkaran 1 (pusat dan jari-jari) dan menyimpannya
	fmt.Scan(&cy1, &cx1, &r1)

	// membaca input untuk lingkaran 2 (pusat dan jari-jari) dan menyimpannya
	fmt.Scan(&cy2, &cx2, &r2)

	// membaca input untuk titik sembarang (koordinat titik) dan menyimpannya
	fmt.Scan(&y, &x)

	// konversi nilai inputan ke float64 agar sinkron dengan fungsi perhitungan
	cy1f, cx1f, r1f := float64(cy1), float64(cx1), float64(r1)
	cy2f, cx2f, r2f := float64(cy2), float64(cx2), float64(r2)
	yf, xf := float64(y), float64(x)

	// mengecek apakah titik berada di dalam lingkaran 1 dan/atau lingkaran 2
	dalam1 := didalam(cx1f, cy1f, r1f, xf, yf)
	dalam2 := didalam(cx2f, cy2f, r2f, xf, yf)

	// output hasil berdasarkan kondisi apakah titik berada di dalam lingkaran 1 dan/atau 2
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

// membuat fungsi untuk menghitung jarak antara dua titik (a,b) dan (c,d)
func jarak(a, b, c, d float64) float64 {
	// menggunakan rumus jarak Euclidean antara dua titik
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// membuat fungsi untuk mengecek apakah titik (x,y) berada di dalam lingkaran dengan pusat (cx,cy) dan jari-jari r
func didalam(cx, cy, r, x, y float64) bool {
	// menghitung jarak dari titik (x, y) ke pusat lingkaran (cx, cy)
	// jika jarak kurang dari atau sama dengan jari-jari lingkaran, maka titik berada di dalam lingkaran
	return jarak(cx, cy, x, y) <= r
}