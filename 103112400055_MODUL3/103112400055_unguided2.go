//Feros Pedrosa Valentino

package main

import "fmt"

func kuadrat(angka int) int {
	return angka * angka
}

func kurangDua(angka int) int {
	return angka - 2
}

func tambahSatu(angka int) int {
	return angka + 1
}

func komposisiFoGoH(angka int) int {
	return kuadrat(kurangDua(tambahSatu(angka)))
}

func komposisiGoHoF(angka int) int {
	return kurangDua(tambahSatu(kuadrat(angka)))
}

func komposisiHoFoG(angka int) int {
	return tambahSatu(kuadrat(kurangDua(angka)))
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(komposisiFoGoH(a))
	fmt.Println(komposisiGoHoF(b))
	fmt.Println(komposisiHoFoG(c))
}
