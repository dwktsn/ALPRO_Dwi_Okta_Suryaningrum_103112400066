package main

import "fmt"

func main() {
	var n int
	var teks string
	fmt.Scan(&n)
	teks = "bukan"
	if n%2 == 0 && n<0 {
		teks = "genap negatif"
	}
	fmt.Println(teks)
}