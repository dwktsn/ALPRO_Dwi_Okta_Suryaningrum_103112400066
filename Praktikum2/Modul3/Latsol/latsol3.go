package main

import "fmt"

func main(){
	var tahun int
	fmt.Print("Masukkan tahun :")
	fmt.Scan(&tahun)
	if tahun % 400 == 0 {
		fmt.Println("Kabisat : True")
	} else if tahun % 4 == 0 {
		fmt.Println("Kabisat : True")
	} else if tahun % 100 == 0{
		fmt.Println("Kabisat : False")
	} else {
		fmt.Println("Kabisat : False")
	}
}