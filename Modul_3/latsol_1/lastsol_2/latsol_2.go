package main

import (
	"fmt"
)

func kuadrat(x int) int {
	return x * x
}

func kurangiDua(x int) int {
	return x - 2
}

func tambahSatu(x int) int {
	return x + 1
}

func hitungFGH(x int) int {
	return kuadrat(kurangiDua(tambahSatu(x)))
}

func hitungGHF(x int) int {
	return kurangiDua(tambahSatu(kuadrat(x)))
}

func hitungHFG(x int) int {
	return tambahSatu(kuadrat(kurangiDua(x)))
}

func main() {

	var nilai1, nilai2, nilai3 int

	fmt.Println("Masukkan tiga bilangan bulat, dipisahkan oleh spasi:")
	fmt.Scan(&nilai1, &nilai2, &nilai3)

	hasil1 := hitungFGH(nilai1)
	hasil2 := hitungGHF(nilai2)
	hasil3 := hitungHFG(nilai3)

	fmt.Println("Hasil dari (f ∘ g ∘ h)(a):", hasil1)
	fmt.Println("Hasil dari (g ∘ h ∘ f)(b):", hasil2)
	fmt.Println("Hasil dari (h ∘ f ∘ g)(c):", hasil3)
}
