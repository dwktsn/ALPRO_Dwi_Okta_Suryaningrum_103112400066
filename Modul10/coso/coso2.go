package main

import "fmt"


func main() {
	var x rune
	var huruf, vkecil, vbesar bool
	fmt.Scanf("%c", &x)
	huruf = (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z')
	vkecil = x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o'
	vbesar = x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O'

	if huruf && (vkecil || vbesar) {
		fmt.Println("vokal")
	}else if huruf && !(vkecil || vbesar){
		fmt.Println("konsonan")
	}else{
		fmt.Println("bukan huruf")
	}
}