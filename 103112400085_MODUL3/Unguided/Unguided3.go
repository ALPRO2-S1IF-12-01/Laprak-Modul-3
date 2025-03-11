// Anastasia Adinda Narendra Indrianto
// 103112400085 S1IF-12-01
package main

import (
	"fmt"
	"math"
)

func hitungJarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func dalamLingkaran(cx, cy, r, x, y int) bool {
	return hitungJarak(x, y, cx, cy) <= float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)

	diLingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
	diLingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
