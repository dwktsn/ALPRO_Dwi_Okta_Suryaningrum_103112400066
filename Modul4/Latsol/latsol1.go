package main

import "fmt"

func main(){
	var belanjaawal, diskon, total int
	fmt.Scan(&belanjaawal)
	fmt.Scan(&diskon)

	total = belanjaawal - (belanjaawal * diskon)/100
	fmt.Println(total)
}
