package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for (a+b)%2 == 0{
		fmt.Println("Hasil penjumlahan", a+b)
		fmt.Scan(&a, &b)
	}
	fmt.Println("program selesai")
}