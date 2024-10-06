package main

import "fmt"

func main(){
	var celcius, fahrenheit, reamur, kelvin int
	fmt.Scan(&celcius)

	fahrenheit = (celcius * 9/5 ) + 32 
	reamur = celcius * 4/5
	kelvin = celcius + 273

	println("Temperatur Celcius :", celcius)
	println("Temperatur Reamur :", reamur)
	println("Temperatur Fahrenheit :", fahrenheit)
	println("Temperatur Kelvin :", kelvin)
}