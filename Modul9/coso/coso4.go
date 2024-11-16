package main

import "fmt"

func main(){
	var bilangan int
	var hasil bool
	fmt.Scan(&bilangan)
	if bilangan < 0 && bilangan%2 == 0 {
		hasil = true
	}
	// hasil = bilangan < 0 && bilangan%2 == 0
	fmt.Println(hasil)
}