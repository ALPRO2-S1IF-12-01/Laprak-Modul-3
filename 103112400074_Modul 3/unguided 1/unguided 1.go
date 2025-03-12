//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
    var a, b, c, d int
    
    fmt.Scan(&a, &b, &c, &d)

    p1:=permutasi(a, c)
    c1 := kombinasi(a, c)
    fmt.Printf("%d %d\n", p1, c1)
    
    p2 :=permutasi(b, d)
    c2 := kombinasi(b, d)
    fmt.Printf("%d %d\n", p2, c2)
}
func faktorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * faktorial(n-1)
}

func permutasi(n, r int) int {
    return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
    return faktorial(n) / (faktorial(r) * faktorial(n-r))
}