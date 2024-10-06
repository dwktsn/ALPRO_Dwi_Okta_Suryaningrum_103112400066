package main

import "fmt"

func main(){
	var  x, operasi float32
	
	fmt. Print("masukkan x : ")
	fmt.Scan(&x)

	operasi = (2/(x+5)+5)
	fmt.Println("Hasil dari : ", x, "adalah ", operasi)
}