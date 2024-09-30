package main

import "fmt"

func main(){
	var mk string = "Algoritma dan Pemrograman"
	var kode string 
	var sks int

	fmt.Print("Tuliskan kode dan sks :")
	fmt.Scan(&kode, &sks)
	fmt.Print("Kredit MK", kode, " - ", mk, "1 adalah", sks, "SKS")
}