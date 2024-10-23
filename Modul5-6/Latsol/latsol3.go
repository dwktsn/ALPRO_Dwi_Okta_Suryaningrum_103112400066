package main

import "fmt"
import "math"

func main(){
	var n, m, hasil float64
	fmt.Scan (&n,&m)
	hasil = math.Pow(n, m)

	fmt.Println(hasil)
}