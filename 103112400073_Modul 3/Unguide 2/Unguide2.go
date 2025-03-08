//Muhammad Zaky Mubarok
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

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    
    f_g_h_a := f(g(h(a)))    
    g_h_f_b := g(h(f(b)))    
    h_f_g_c := h(f(g(c)))    

    
    fmt.Println( f_g_h_a)
    fmt.Println( g_h_f_b)
    fmt.Println( h_f_g_c)
}
