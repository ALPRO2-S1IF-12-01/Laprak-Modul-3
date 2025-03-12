//Achmad Zulvan Nur Hakim 103112400070
package main

import (
	"fmt"
	"math"
)

func jarak(n1, n2,n3,n4 float64) float64 {
	return math.Sqrt(math.Pow(n1-n3, 2) + math.Pow(n2-n4, 2))
}

func Titik(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, cx2, cy2, r1, r2 float64
	fmt.Scan(&cx1, &cy1, &cx2, &cy2, &r1, &r2)
	
	var x, y float64
	fmt.Scan(&x, &y)

	Titik1 := Titik(cx1, cy1, r1, x, y)
	Titik2 := Titik(cx2, cy2, r2, x, y)

	if Titik1 && Titik2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if Titik1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if Titik2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}