package main

import "fmt"

func main(){
	var nama, nim, kelas string
	fmt.Print("Masukkan nama anda : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan kelas anda : ")
	fmt.Scan(&kelas)
	fmt.Print("Masukkan nim anda : ")
	fmt.Scan(&nim)

	fmt.Println ("Perkenalkan saya adalah " + nama + " salah satu mahasiswa Prodi S1-IF dari kelas " + kelas + " dengan NIM " + nim)

}