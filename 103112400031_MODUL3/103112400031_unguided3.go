// Savila Nur Fadilla
// 103112400031

package main

import (
	"fmt"
	"math"
)

func hitungJarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func inSide(cx, cy, r, x, y int) bool {
	return hitungJarak(float64(cx), float64(cy), float64(x), float64(y)) <= float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	inCircle1 := inSide(cx1, cy1, r1, x, y)
	inCircle2 := inSide(cx2, cy2, r2, x, y)

	if inCircle1 && inCircle2 {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}