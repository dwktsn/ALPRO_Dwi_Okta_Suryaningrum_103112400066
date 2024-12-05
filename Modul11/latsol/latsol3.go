package main

import "fmt"

func main() {
	var input, hasil int
	fmt.Scan(&input)

	switch{
	case input%10 == 0:
		fmt.Println("kategori = Bilangan Kelipatan 10")
		hasil = input/10
		fmt.Printf("Hasil pembagian antara (%d / 10) = %d\n", input, hasil)
	case input%5 == 0 && input != 5:
		fmt.Println("kategori = Bilangan Kelipatan 5")
		hasil = input*input
		fmt.Printf("Hasil kuadrat dari  (%d ^ 2) = %d\n", input, hasil)
	case input%2 == 0:
		hasil = input*(input+1)
		fmt.Println("kategori = Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya (%d * %d) = %d\n", input, input+1, hasil)
	case input%2 != 0:
		hasil = input+(input+1)
		fmt.Println("kategori = Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya (%d + %d) = %d\n", input, input+1, hasil)
	}
}