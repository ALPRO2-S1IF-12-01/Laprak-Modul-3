//ABID FADHILAH M
//103112400046
package main

import "fmt"

func main() {
    var a, d, A int
    fmt.Scan(&a, &d, &A)
    fmt.Println(fogoh(a))
    fmt.Println(gohof(d))
    fmt.Println(hofog(A))
}

func f(x int) int { return x * x
}

func g(x int) int { return x - 2
}

func h(x int) int { return x + 1
}

func fogoh(x int) int { return f(g(h(x)))
}

func gohof(x int) int { return g(h(f(x)))
}

func hofog(x int) int { return h(f(g(x)))
}