package main

import "fmt"

func main() {
	var word string
	var repeat int
	fmt.Scan(&word, &repeat)
	counter := 0
	for done := false; !done;{
		fmt.Println(word)
		counter++
		done = (counter>=repeat)
	}
}

// for i := 0; i < repeat; i++ {
	// 	fmt.Println(word)
	// }