package main

import (
    "fmt"
)
func kuadrat(x int) int {
    return x * x
}
func kurangDua(x int) int {
    return x - 2
}
func tambahSatu(x int) int {
    return x + 1
}
func komposisiFoGoH(x int) int {
    return kuadrat(kurangDua(tambahSatu(x)))
}
func komposisiGoHoF(x int) int {
    return kurangDua(tambahSatu(kuadrat(x)))
}
func komposisiHoFoG(x int) int {
    return tambahSatu(kuadrat(kurangDua(x)))
}
func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    fmt.Println(komposisiFoGoH(a))
    fmt.Println(komposisiGoHoF(b))
    fmt.Println(komposisiHoFoG(c))
}
