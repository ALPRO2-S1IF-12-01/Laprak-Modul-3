//MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"

func main() {
        var a, b, c int
        fmt.Scan(&a, &b, &c)
        fmt.Println(fogoh(a))
        fmt.Println(gohof(b))
        fmt.Println(hofog(c))
}

func f(x int) int {
        return x * x
}

func g(x int) int {
        return x - 2
}

func h(x int) int {
        return x + 1
}

func fogoh(x int) int {
        return f(g(h(x)))
}

func gohof(x int) int {
        return g(h(f(x)))
}

func hofog(x int) int {
        return h(f(g(x)))
}