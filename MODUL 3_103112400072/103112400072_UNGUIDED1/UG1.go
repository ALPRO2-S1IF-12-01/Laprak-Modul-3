// MUHAMAD FAZA FAHRI AZIZ || 103112400072

package main

import "fmt"

func main() {
	var a, b,c,d int
	fmt.Scan(&a, &b,&c,&d)
	

	fmt.Print(permutasi(a,c)," ",kombinasi(a,c)," ")
	fmt.Print(permutasi(b,d)," ", kombinasi(b,d))

}
func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	if r > n {
		return 0 
	}
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n,r int)int  {
	if r > n {
		return 0 
	}
	return faktorial(n) / (faktorial(r) * (faktorial(n - r)))
	
}