// PRATAMA BINTANG DANISWARA 103112400051
package main

import "fmt"

func dl(pX, pY, r, tX, tY int) bool {
	j := (tX-pX)*(tX-pX) + (tY-pY)*(tY-pY)
	return j <= r*r
}
func main() {
	var pX1, pY1, r1 int
	var pX2, pY2, r2 int
	var tX, tY int
	fmt.Scan(&pX1, &pY1, &r1)
	fmt.Scan(&pX2, &pY2, &r2)
	fmt.Scan(&tX, &tY)
	dl1 := dl(pX1, pY1, r1, tX, tY)
	dl2 := dl(pX2, pY2, r2, tX, tY)
	if dl1 && dl2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dl1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dl2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
