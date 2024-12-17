package main

import "fmt"

func main() {
	var n int
	var result bool
	for result = true; result;{
		fmt.Scan(&n)
		result = n <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", n)
}