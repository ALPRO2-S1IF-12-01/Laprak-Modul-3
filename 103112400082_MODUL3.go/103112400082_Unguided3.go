//RizkinaAzizah
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

	fmt.Println(posisi(cx1, cy1, r1, cx2, cy2, r2, x, y))
}

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func didalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func posisi(cx1, cy1, r1, cx2, cy2, r2, x, y int) string {
	square1 := didalam(cx1, cy1, r1, x, y)
	square2 := didalam(cx2, cy2, r2, x, y)

	if square1 && square2 {
		return "Titik didalam lingkaran 1 dan 2"
	} else if square1 {
		return "Titik didalam lingkaran 1"
	} else if square2 {
		return "Titik didalam lingkaran 2"
	} else {
		return "Titik diluar lingkaran 1 dan 2"
	}
}
