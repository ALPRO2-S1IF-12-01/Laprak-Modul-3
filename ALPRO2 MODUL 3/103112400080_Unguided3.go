// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func dalamLingkaran(cx, cy, r, x, y int) bool {
	jarak := (x-cx)*(x-cx) + (y-cy)*(y-cy)
	return jarak <= r*r
}

func main() {
	var x1, y1, r1, x2, y2, r2, x, y int
	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)

	in1 := dalamLingkaran(x1, y1, r1, x, y)
	in2 := dalamLingkaran(x2, y2, r2, x, y)

	switch {
	case in1 && in2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case in1:
		fmt.Println("Titik di dalam lingkaran 1")
	case in2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
