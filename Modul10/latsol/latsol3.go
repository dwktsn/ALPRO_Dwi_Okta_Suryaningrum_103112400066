package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}

	if b%2 == 0 || b%3 == 0 || b%5 == 0 || b%7 == 0 && b != 1 && b != 2 && b != 3 && b != 5 && b != 7 {
		fmt.Println("\nFalse")
	} else {
		fmt.Println("\nTrue")
	}
}