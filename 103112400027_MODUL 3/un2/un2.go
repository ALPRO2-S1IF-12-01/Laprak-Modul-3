//Raja Muhammad Lufhti 103112400027 
package main

import "fmt"

func main() {
	var k, l, m int
	fmt.Scan(&k, &l, &m)
	fmt.Println("fogog:", fogog(k))
	fmt.Println("gohof:", gohof(l))
	fmt.Println("hofog:", hofog(m))
}

func a(x int) int {
        return x * x
}

func b(x int) int {
        return x - 2
}

func c(x int) int {
        return x + 1
}

func fogog(x int) int {
        return a(b(c(x)))
}

func gohof(x int) int {
        return b(c(a(x)))
}

func hofog(x int) int {
        return c(a(b(x)))
}
