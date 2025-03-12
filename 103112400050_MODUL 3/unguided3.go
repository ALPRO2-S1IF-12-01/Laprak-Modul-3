package main

import (
	"fmt"
	"math"
	//ARIEL AHNAF KUSUMA 103112400050
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func main() {
	var x1, y1, r1, x2, y2, r2, x, y float64
	fmt.Scan(&x1, &y1, &r1, &x2, &y2, &r2, &x, &y)

	jarak1 := jarak(x1, y1, x, y)
	jarak2 := jarak(x2, y2, x, y)

	if jarak1 <= r1 && jarak2 <= r2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if jarak1 <= r1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if jarak2 <= r2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
