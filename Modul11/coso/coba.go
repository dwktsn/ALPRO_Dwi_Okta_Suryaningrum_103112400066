package main

import "fmt"

func main() {
	var input, hasil, nomor int
	var kategori, deskripsi string

	fmt.Print("Masukkan nomor masukan: ")
	fmt.Scan(&nomor)
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&input)

	switch {
	case input%10 == 0: // Kelipatan 10
		kategori = "Bilangan Kelipatan 10"
		hasil = input / 10
		deskripsi = fmt.Sprintf("Hasil pembagian antara %d / 10 = %d", input, hasil)
	case input%5 == 0: // Kelipatan 5 (bukan kelipatan 10)
		kategori = "Bilangan Kelipatan 5"
		hasil = input * input
		deskripsi = fmt.Sprintf("Hasil kuadrat dari %d ^ 2 = %d", input, hasil)
	case input%2 == 0: // Bilangan Genap
		kategori = "Bilangan Genap"
		hasil = input * (input + 1)
		deskripsi = fmt.Sprintf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d", input, input+1, hasil)
	case input%2 != 0: // Bilangan Ganjil
		kategori = "Bilangan Ganjil"
		hasil = input + (input + 1)
		deskripsi = fmt.Sprintf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", input, input+1, hasil)
	default:
		kategori = "Tidak Termasuk Kategori"
		deskripsi = "Bilangan tidak memenuhi salah satu kategori."
	}

	// Output sesuai format
	fmt.Printf("No Masukan: %d\n", nomor)
	fmt.Printf("Kategori: %s\n", kategori)
	fmt.Println(deskripsi)
}
