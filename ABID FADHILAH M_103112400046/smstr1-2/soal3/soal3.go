//ABID FADHILAH M
//103112400046
package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x1-x2, y1-y2)
}

func main() {
	var x1, y1, r1, x2, y2, r2, x, y float64
	fmt.Scan(&x1, &y1, &r1, &x2, &y2, &r2, &x, &y)

	j1, j2 := jarak(x1, y1, x, y), jarak(x2, y2, x, y)
	switch {
	case j1 <= r1 && j2 <= r2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case j1 <= r1:
		fmt.Println("Titik di dalam lingkaran 1")
	case j2 <= r2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
