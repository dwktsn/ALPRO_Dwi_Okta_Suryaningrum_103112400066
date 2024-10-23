package main

import "fmt"

func main(){
	var a,b,j int
	fmt.Scan(&a, &b)
	for j = a; j <= b; j=j+1{
		fmt.Print(j, "")
	}
}