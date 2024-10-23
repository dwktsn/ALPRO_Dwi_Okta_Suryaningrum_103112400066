package main

import "fmt"
import "math"

func main() {
	var n int
	var volume, jarialas, tinggi float64
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&jarialas, &tinggi)
		volume = (1.0/3.0) * (math.Pi * jarialas* jarialas * tinggi) 
		fmt.Println(volume)
	}
}
