package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y float64
	)

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

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

func jarak(a, b, c, d float64) float64{
	jarak := math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
	return jarak
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}