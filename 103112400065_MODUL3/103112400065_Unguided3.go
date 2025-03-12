// Dimas Ramadhani
// 103112400065

package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	Lingkaran1:= didalam(cx1, cy1, r1, x, y)
	Lingkaran2:= didalam(cx2, cy2, r2, x, y)

	if Lingkaran1 && Lingkaran2 {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if Lingkaran1 {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if Lingkaran2 {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}

func didalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(math.Pow(float64(a-c), 2) + math.Pow(float64(b-d), 2))
}	