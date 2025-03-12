//Ahmad Ruba'i
// 103112400074
package main

import (
    "fmt"
    "math"
)

func jarak(a, b, c, d float64) float64 {
    return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
    return jarak(cx, cy, x, y) <= r
}

func main() {
    var x1, y1, r1 float64 
    var x2, y2, r2 float64 
    var x, y float64       

    fmt.Scan(&x1, &y1, &r1)
    fmt.Scan(&x2, &y2, &r2)
    fmt.Scan(&x, &y)

    dalamL1 := didalam(x1, y1, r1, x, y)
    dalamL2 := didalam(x2, y2, r2, x, y)

    if dalamL1 && dalamL2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if dalamL1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if dalamL2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}