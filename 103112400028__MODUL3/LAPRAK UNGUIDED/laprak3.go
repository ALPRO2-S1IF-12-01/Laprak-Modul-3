// MUHAAMAD GAMEL AL GHIFARI
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func dalmLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1 float64
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 float64
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y float64
	fmt.Scan(&x, &y)

	dalmLingkaran1 := dalmLingkaran(cx1, cy1, r1, x, y)
	dalmLingkaran2 := dalmLingkaran(cx2, cy2, r2, x, y)

	if dalmLingkaran1 && dalmLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalmLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalmLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}