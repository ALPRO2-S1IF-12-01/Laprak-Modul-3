//BERTHA ADELA
//103112400041
package main
import (
	"fmt"
	"math"
)
func main() {
	var cx1, cy1, r1, r2, cx2, cy2, x, y float64
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)
	if Didalam(cx1, cy1, r1, x, y) && Didalam(cx2, cy2, r2, x, y) {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if Didalam(cx1, cy1, r1, x, y) {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if Didalam(cx2, cy2, r2, x, y) {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}
func Jarak(a, b, c, d float64) float64 {
	return math.Sqrt((c-a)*(c-a)+(d-b)*(d-b))
}
func Didalam(cx, cy, r, x, y float64) bool {
	Jarak := Jarak(cx, cy, x, y)
	return Jarak <= r
}