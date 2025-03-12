package main
// 103112400078
// Mohammad Reyhan Aretha Fatin
import "fmt"

func faktorial(n int) int {
	if n <= 1 {
		return 1
	}
	fktr := 1
	for i := 2; i <= n; i++ {
		fktr *= i
	}
	return fktr
}

func hitungPermutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func hitungKombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println(hitungPermutasi(a, c), hitungKombinasi(a, c))
		fmt.Println(hitungPermutasi(b, d), hitungKombinasi(b, d))
	} else {
		fmt.Println("Input tidak sesuai")
	}
}
