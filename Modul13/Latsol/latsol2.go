package main

import "fmt"

func main() {

	var angka, done float32

	fmt.Scan(&angka)

	done = float32(int(angka)) + 1

	for selesai := false; !selesai; {
		angka += 0.1
		angka = float32(int(angka*10+0.5)) / 10

		if angka >= done {
			fmt.Printf("%.0f\n", angka)
			selesai = true
		} else {
			fmt.Printf("%.1f\n", angka)
		}

	}

}