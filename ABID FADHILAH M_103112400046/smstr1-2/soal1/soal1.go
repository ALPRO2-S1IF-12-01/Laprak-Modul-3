//ABID FADHILAH M
//103112400046
package main

import "fmt"

func main() {
    var a, b, i, d int
    fmt.Scan(&a, &b, &i, &d)

    if a >= i && b >= d {
        fmt.Println(hitungPermutasi(a, i), hitungKombinasi(a, i))
        fmt.Println(hitungPermutasi(b, d), hitungKombinasi(b, d))
    } else {
        fmt.Println("Input tidak sesuai")
    }
}

func faktorial(n int) int {
    hasil := 1
    for j := 1; j <= n; j++ {
        hasil *= j
    }
    return hasil
}

func hitungPermutasi(n, r int) int {
    if r > n {
        return 0
    }
    return faktorial(n) / faktorial(n-r)
}

func hitungKombinasi(n, r int) int {
    if r > n {
        return 0
    }
    return faktorial(n) / (faktorial(r) * faktorial(n-r))
}