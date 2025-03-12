//Raja Muhammad Lufhti 103112400027
package main

import (
	"fmt"
	"math"
)

func jarak(p, b, m, n float64) float64 {
	return math.Sqrt(math.Pow(float64(m-p), 2) + math.Pow(float64(n-b), 2))
}

func dalam(xy, yx, b, x, y float64) bool {
	return jarak(xy, yx, x, y) <= float64(b) 
}

func main() {
	var m1, P1, d1 float64
	fmt.Scan(&m1, &P1, &d1)

	var m2, P2, d2 float64
	fmt.Scan(&m2, &P2, &d2)

	var x, y float64
	fmt.Scan(&x, &y)

	dalamLingkaran1 := dalam(m1, P1, d1, x, y)
	dalamLingkaran2 := dalam(m2, P2, d2, x, y)

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