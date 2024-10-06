package main

import "fmt"
import "math"

func main(){
	var jarijari, volumebola, luasbola float64
	fmt.Scan(&jarijari)
	
	volumebola = (4.0/3.0) * math.Pi * math.Pow(jarijari, 3)
	luasbola = 4 * math.Pi * math.Pow(jarijari, 2)

	fmt.Printf("Bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f\n", jarijari, volumebola, luasbola)
}