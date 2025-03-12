//Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func f(x int) int {
        return x * x
}

func g(x int) int {
        return x - 2
}

func h(x int) int {
        return x + 1
}

func fOgOh(x int) int {
        return f(g(h(x)))
}

func gOhOf(x int) int {
        return g(h(f(x)))
}

func hOfOg(x int) int {
        return h(f(g(x)))
}
func main() {
	var p, q, r int
	fmt.Scan(&p, &q, &r)
	fmt.Println(fOgOh(p))
	fmt.Println(gOhOf(q))
	fmt.Println(hOfOg(r))
}