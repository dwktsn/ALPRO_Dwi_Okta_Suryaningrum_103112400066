package main

import "fmt"

func main(){
	var f, c int
	fmt.Scan(&f)
	c = (f - 32 ) * 5/9 
	fmt.Println(c)
}