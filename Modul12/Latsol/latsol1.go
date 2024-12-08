package main

import "fmt"

func main() {
	var usr, pwd string
	var salah int

	salah = 0
	fmt.Scan(&usr, &pwd)
	for usr != "Admin" || pwd != "Admin"{
		fmt.Scan(&usr, &pwd)
		salah++
	}
	fmt.Printf("%d percobaan gagal login\n", salah)
}
