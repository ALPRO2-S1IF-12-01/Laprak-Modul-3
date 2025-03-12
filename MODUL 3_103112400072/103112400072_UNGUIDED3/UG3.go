package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func didalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func cekPosisi(cx1, cy1, r1, cx2, cy2, r2, x, y float64) string {

	dalam1 := didalamLingkaran(cx1, cy1, r1, x, y)
	dalam2 := didalamLingkaran(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dalam1 {
		return "Titik di dalam lingkaran 1"
	} else if dalam2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {

	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64

	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Scan(&cx2, &cy2, &r2)

	var x, y float64
	fmt.Scan(&x, &y)

	hasil := cekPosisi(cx1, cy1, r1, cx2, cy2, r2, x, y)
	fmt.Println(hasil)
}
