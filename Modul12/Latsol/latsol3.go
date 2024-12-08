package main

import "fmt"

func main() {
	var x, y int
	
	fmt.Scan(&x, &y)

	i := 0
	temp := x
	for temp >= y{
		temp -=y
		i++
	}
	fmt.Println(i)
}