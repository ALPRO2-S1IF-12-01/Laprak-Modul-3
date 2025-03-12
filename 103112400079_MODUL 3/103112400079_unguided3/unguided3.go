package main

// Muhammad Faris Rachmadi
// 103112400079

import "fmt"

func DalamLingkaran(cx, cy, r, x, y float64) bool {
	jarakKuadrat := (x-cx)*(x-cx) + (y-cy)*(y-cy)
	return jarakKuadrat <= r*r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	diDalamLingkaran1 := DalamLingkaran(cx1, cy1, r1, x, y)
	diDalamLingkaran2 := DalamLingkaran(cx2, cy2, r2, x, y)

	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
