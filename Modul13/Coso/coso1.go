package main

import "fmt"

func main() {
	var usr, pw string
	var kondisi bool

	for kondisi = false; !kondisi;  {
		fmt.Scan(&usr,&pw)
		kondisi = usr == "admin" && pw == "admin12345"
	}
	fmt.Println("selamat anda berhasil login")
}