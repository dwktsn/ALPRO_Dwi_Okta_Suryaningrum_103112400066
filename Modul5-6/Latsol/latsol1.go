package main

import "fmt"

func main(){
	var n, hasil int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hasil += i
	}
	fmt.Printf("%d\n", hasil)
}