//Nama : Anggun Wahyu W. (103112480280)
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) < r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	fCx1, fCy1, fR1 := float64(cx1), float64(cy1), float64(r1)
	fCx2, fCy2, fR2 := float64(cx2), float64(cy2), float64(r2)
	fX, fY := float64(x), float64(y)

	dalamLingkaran1 := didalam(fCx1, fCy1, fR1, fX, fY)
	dalamLingkaran2 := didalam(fCx2, fCy2, fR2, fX, fY)

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