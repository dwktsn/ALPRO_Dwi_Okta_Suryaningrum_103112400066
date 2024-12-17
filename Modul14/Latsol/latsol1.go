package main

import "fmt"

func main() {
	var n int
	var result string

	fmt.Scan(&n)
	result = "prima"

	if n%n == 0 && n%2 != 0 {
		fmt.Println(result)
	}else{
		fmt.Println("Bukan Prima")
	}
}