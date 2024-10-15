package main

import "fmt"
import "math"

func main(){
	var x1,x2,x3,y1,y2,y3 float64
	fmt.Scan(&x1,&y1)
	fmt.Scan(&x2,&y2)
	fmt.Scan(&x3,&y3)

	panjangAB := math.Sqrt(math.Pow(x2-x1,2)+math.Pow(y2-y1, 2))
	panjangBC := math.Sqrt(math.Pow(x3-x2,2)+math.Pow(y3-y2, 2))
	panjangCA := math.Sqrt(math.Pow(x1-x3,2)+math.Pow(y1-y3, 2))

	panjangTerpanjang := math.Max(panjangAB, math.Max(panjangBC, panjangCA))

	fmt.Println(panjangTerpanjang)
}