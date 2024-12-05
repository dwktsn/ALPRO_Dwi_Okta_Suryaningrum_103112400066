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
	case tipe_kendaraan == "motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case tipe_kendaraan == "motor" && durasi > 2:
		tarif = 9000
	case tipe_kendaraan == "mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case tipe_kendaraan == "mobil" && durasi > 2:
		tarif = 20000
	case tipe_kendaraan == "truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case tipe_kendaraan == "truk" && durasi > 2:
		tarif = 30000
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif parkir : Rp %d\n", tarif)
}