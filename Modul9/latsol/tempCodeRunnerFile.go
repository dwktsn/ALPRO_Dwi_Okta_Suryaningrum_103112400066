package main

import "fmt"

func main() {
	var x,y int
	var result bool
	fmt.Scan(&x, &y)
	if y%x == 0 {
		result = true
	}else{
		result = false
	}
	fmt.Println(result)

	if x%y == 0 {
		result = true
	}else{
		result = false
	}
	fmt.Println(result)
}