// Felix Pedrosa V 

package main

import "fmt"

// Fungsi f(x) = x^2
func f(x int) int {
    return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
    return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
    return x + 1
}

// Komposisi f(g(h(x)))
func fogoh(x int) int {
    return f(g(h(x)))
}

// Komposisi g(h(f(x)))
func gohof(x int) int {
    return g(h(f(x)))
}

// Komposisi h(f(g(x)))
func hofog(x int) int {
    return h(f(g(x)))
}

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    fmt.Println(fogoh(a))
    fmt.Println(gohof(b))
    fmt.Println(hofog(c))
}