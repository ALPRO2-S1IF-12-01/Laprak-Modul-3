package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func didalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	// mengecek titik di lingkrn 1 ap 2
	inside1 := didalam(cx1, cy1, r1, x, y)
	inside2 := didalam(cx2, cy2, r2, x, y)

	if inside1 && inside2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inside1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inside2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
