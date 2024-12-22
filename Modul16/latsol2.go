package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Set seed untuk memastikan hasil random tetap konsisten
	rand.Seed(1)

	// Input jumlah tetesan hujan
	var drops int
	fmt.Scan(&drops)

	// Hitung curah hujan di setiap daerah
	const dropSize = 0.0001 // Ukuran tiap tetesan dalam ml
	countA, countB, countC, countD := 0, 0, 0, 0

	for i := 0; i < drops; i++ {
		x, y := rand.Float64(), rand.Float64()

		if x < 0.5 && y >= 0.5 {
			countA++
		} else if x >= 0.5 && y >= 0.5 {
			countB++
		} else if x < 0.5 && y < 0.5 {
			countC++
		} else if x >= 0.5 && y < 0.5 {
			countD++
		}
	}

	// Hitung curah hujan
	rainA := float64(countA) * dropSize
	rainB := float64(countB) * dropSize
	rainC := float64(countC) * dropSize
	rainD := float64(countD) * dropSize

	// Output hasil
	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", rainA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", rainB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", rainC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", rainD)
}
