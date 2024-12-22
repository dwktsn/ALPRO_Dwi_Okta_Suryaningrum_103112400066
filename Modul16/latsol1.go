package main

import "fmt"

func main() {
	var number, sum float64
	var count int

	for {
		fmt.Scan(&number)

		if number == 9999 {
			break
		}

		sum += number
		count++
	}

	if count == 0 {
		fmt.Println("Tidak ada bilangan untuk dihitung rata-ratanya.")
	} else {
		average := sum / float64(count)
		fmt.Printf("Rata-rata : %.2f\n", average)
	}
}
