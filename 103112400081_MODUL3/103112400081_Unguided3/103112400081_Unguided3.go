// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func checkPosition(cx1, cy1, r1, cx2, cy2, r2, x, y float64) string {
	dist1 := jarak(cx1, cy1, x, y)
	dist2 := jarak(cx2, cy2, x, y)
	if dist1 <= r1 && dist2 <= r2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dist1 <= r1 {
		return "Titik di dalam lingkaran 1"
	} else if dist2 <= r2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di dalam lingkaran 1 dan 2"
	}
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64
	fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)
	fmt.Println(checkPosition(cx1, cy1, r1, cx2, cy2, r2, x, y))
}
