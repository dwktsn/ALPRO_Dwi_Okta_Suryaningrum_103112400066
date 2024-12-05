package main

import "fmt"

func main() {
	var tipe_kendaraan string
	var durasi, tarif int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk) : ")
	fmt.Scan(&tipe_kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam) : ")
	fmt.Scan(&durasi)

	switch {
	case tipe_kendaraan == "motor":
		tarif = 2000*durasi
	case tipe_kendaraan == "mobil":
		tarif = 5000*durasi
	case tipe_kendaraan == "truk":
		tarif = 8000*durasi
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Println("Rp", tarif)
}